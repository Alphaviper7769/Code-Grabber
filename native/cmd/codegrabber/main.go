package main

import (
	"native/internal/application"
	"native/internal/domain"
	"native/internal/infrastructure/config"
	"native/internal/infrastructure/factory"
	"native/internal/infrastructure/logger"
	"native/internal/infrastructure/native"
)

func main() {

	er := logger.Init("codegrabber.log")
	if er != nil {
		return
	}

	logger.Logger.Println("Native host started")

	// 🔥 VERY IMPORTANT
	cfg, err := config.Load("config.json")
	if err != nil {
		logger.Logger.Println(err)
	}
	// Prevent log pollution of stdout (native protocol)
	// All logging now uses logger.Logger

	logger.Logger.Println("CodeGrabber Native Host Started")

	transport := native.NewTransport()

	processorFactory := factory.NewLanguageProcessorFactory(cfg.BaseDir)
	service := application.NewProblemService(processorFactory)

	for {
		var problem domain.Problem

		// 1️⃣ Read message from Chrome
		err := transport.ReadMessage(&problem)
		if err != nil {
			logger.Logger.Println("Read error:", err)
			return
		}

		logger.Logger.Println("Received problem:", problem.Slug)

		// 2️⃣ Process problem (create files etc.)
		err = service.Handle(problem)
		if err != nil {
			logger.Logger.Println("Processing error:", err)

			// Send failure response
			transport.WriteMessage(map[string]string{
				"status":  "error",
				"message": err.Error(),
			})

			continue
		}

		logger.Logger.Println("Finished processing:", problem.Slug)

		// 3️⃣ Send success response
		err = transport.WriteMessage(map[string]string{
			"status": "ok",
		})
		if err != nil {
			logger.Logger.Println("Write error:", err)
			return
		}
	}
}
