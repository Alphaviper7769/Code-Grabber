package main

import (
	"log"
	"native/internal/application"
	"native/internal/domain"
	"native/internal/infrastructure/config"
	"native/internal/infrastructure/factory"
	"native/internal/infrastructure/native"
	"os"
)

func main() {

	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}
	// 🔥 VERY IMPORTANT
	// Prevent log pollution of stdout (native protocol)
	log.SetOutput(os.Stderr)

	log.Println("CodeGrabber Native Host Started")

	transport := native.NewTransport()

	processorFactory := factory.NewLanguageProcessorFactory(cfg.BaseDir)
	service := application.NewProblemService(processorFactory)

	for {
		var problem domain.Problem

		// 1️⃣ Read message from Chrome
		err := transport.ReadMessage(&problem)
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		log.Println("Received problem:", problem.Slug)

		// 2️⃣ Process problem (create files etc.)
		err = service.Handle(problem)
		if err != nil {
			log.Println("Processing error:", err)

			// Send failure response
			transport.WriteMessage(map[string]string{
				"status":  "error",
				"message": err.Error(),
			})

			continue
		}

		log.Println("Finished processing:", problem.Slug)

		// 3️⃣ Send success response
		err = transport.WriteMessage(map[string]string{
			"status": "ok",
		})
		if err != nil {
			log.Println("Write error:", err)
			return
		}
	}
}
