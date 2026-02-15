package application

import "native/internal/domain"

type ProblemService struct {
	factory domain.ProcessorFactory
}

func NewProblemService(factory domain.ProcessorFactory) *ProblemService {
	return &ProblemService{
		factory: factory,
	}
}

func (s *ProblemService) Handle(problem domain.Problem) error {
	processor, err := s.factory.GetProcessor(problem.Language)
	if err != nil {
		return err
	}

	return processor.Process(problem)
}
