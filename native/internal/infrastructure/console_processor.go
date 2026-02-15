package infrastructure

import (
	"fmt"
	"native/internal/domain"
	"os"
)

type ConsoleProcessor struct{}

func NewConsoleProcessor() *ConsoleProcessor {
	return &ConsoleProcessor{}
}

func (c *ConsoleProcessor) Process(problem domain.Problem) error {
	fmt.Fprintln(os.Stderr, "Received Problem:")
	fmt.Fprintln(os.Stderr, "Slug:", problem.Slug)
	fmt.Fprintln(os.Stderr, "Language:", problem.Language)
	fmt.Fprintln(os.Stderr, "Tests:", len(problem.Tests))
	return nil
}
