package processors

import (
	"fmt"
	"native/internal/domain"
	"native/internal/infrastructure/logger"
	"os"
	"path/filepath"
)

type PythonProcessor struct {
	baseDir string
}

func NewPythonProcessor(baseDir string) *PythonProcessor {
	return &PythonProcessor{baseDir: baseDir}
}

func (p *PythonProcessor) Process(problem domain.Problem) error {
	logger.Logger.Println("PythonProcessor: Processing problem:", problem.Slug)
	dir := filepath.Join(
		p.baseDir,
		problem.Source,
		"py",
	)

	logger.Logger.Println("PythonProcessor: Creating directory:", dir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.Logger.Println("PythonProcessor: Directory creation error:", err)
		return err
	}

	filePath := filepath.Join(dir, problem.Slug+".py")
	logger.Logger.Println("PythonProcessor: Target file path:", filePath)

	// Check if file exists
	if _, err := os.Stat(filePath); err == nil {
		// File exists, log and skip creation
		logger.Logger.Println("File already exists for language py:", filePath)
		return nil
	} else if !os.IsNotExist(err) {
		// Some other error
		logger.Logger.Println("PythonProcessor: File stat error:", err)
		return err
	}

	template := generatePythonTemplate(problem)
	logger.Logger.Println("PythonProcessor: Writing template to file.")

	return os.WriteFile(filePath, []byte(template), 0644)
}

func generatePythonTemplate(problem domain.Problem) string {
	return fmt.Sprintf(
		"# %s\n"+
			"# %s\n"+
			"\n"+
			"class Solution:\n"+
			"    pass\n",
		problem.Title,
		problem.URL,
	)
}
