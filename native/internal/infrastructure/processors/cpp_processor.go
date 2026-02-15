package processors

import (
	"fmt"
	"os"
	"path/filepath"

	"native/internal/domain"
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

	dir := filepath.Join(
		c.baseDir,
		problem.Source,
		"cpp",
	)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(dir, problem.Slug+".cpp")

	template := generateCppTemplate(problem)

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
