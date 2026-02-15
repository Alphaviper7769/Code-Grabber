package processors

import (
	"fmt"
	"native/internal/domain"
	"os"
)

type PythonProcessor struct{}

func NewPythonProcessor() *PythonProcessor {
	return &PythonProcessor{}
}

func (p *PythonProcessor) Process(problem domain.Problem) error {
	fmt.Fprintln(os.Stderr, "Generating Python template for:", problem.Slug)
	return nil
}
