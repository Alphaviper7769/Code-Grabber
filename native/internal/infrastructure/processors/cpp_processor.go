package processors

import (
	"fmt"
	"os"
	"path/filepath"

	"native/internal/domain"
	"native/internal/infrastructure/logger"
)

type CppProcessor struct {
	baseDir string
}

func NewCppProcessor(baseDir string) *CppProcessor {
	return &CppProcessor{
		baseDir: baseDir,
	}
}

func (c *CppProcessor) Process(problem domain.Problem) error {

	logger.Logger.Println("CppProcessor: Processing problem:", problem.Slug)
	dir := filepath.Join(
		c.baseDir,
		problem.Source,
		"cpp",
	)

	logger.Logger.Println("CppProcessor: Creating directory:", dir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.Logger.Println("CppProcessor: Directory creation error:", err)
		return err
	}

	filePath := filepath.Join(dir, problem.Slug+".cpp")
	logger.Logger.Println("CppProcessor: Target file path:", filePath)

	// Check if file exists
	if _, err := os.Stat(filePath); err == nil {
		// File exists, log and skip creation
		logger.Logger.Println("File already exists for language cpp:", filePath)
		return nil
	} else if !os.IsNotExist(err) {
		// Some other error
		logger.Logger.Println("CppProcessor: File stat error:", err)
		return err
	}

	template := generateCppTemplate(problem)
	logger.Logger.Println("CppProcessor: Writing template to file.")

	return os.WriteFile(filePath, []byte(template), 0644)
}

func generateCppTemplate(problem domain.Problem) string {
	return fmt.Sprintf(`// %s
// %s

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    
};

int main() {
    return 0;
}
`, problem.Title, problem.URL)
}
