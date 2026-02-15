package application

import (
	"native/internal/domain"
	"native/internal/infrastructure/logger"
)

type ProblemService struct {
	factory domain.ProcessorFactory
}

func NewProblemService(factory domain.ProcessorFactory) *ProblemService {
	return &ProblemService{
		factory: factory,
	}
}

func (s *ProblemService) Handle(problem domain.Problem) error {
	logger.Logger.Println("Handling problem:", problem.Slug, "Language:", problem.Language)
	processor, err := s.factory.GetProcessor(problem.Language)
	if err != nil {
		logger.Logger.Println("Error getting processor:", err)
		return err
	}
	logger.Logger.Println("Processing problem with processor:", problem.Language)
	return processor.Process(problem)
}
