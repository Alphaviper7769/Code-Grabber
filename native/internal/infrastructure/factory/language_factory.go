package factory

import (
	"fmt"
	"native/internal/domain"
	"native/internal/infrastructure/logger"
	"native/internal/infrastructure/processors"
)

type LanguageProcessorFactory struct {
	processors map[string]domain.Processor
}

func NewLanguageProcessorFactory(baseDir string) *LanguageProcessorFactory {

	logger.Logger.Println("Initializing language processor factory with baseDir:", baseDir)
	return &LanguageProcessorFactory{
		processors: map[string]domain.Processor{
			"cpp": processors.NewCppProcessor(baseDir),
			"py":  processors.NewPythonProcessor(baseDir),
		},
	}
}

func (f *LanguageProcessorFactory) GetProcessor(lang string) (domain.Processor, error) {

	logger.Logger.Println("Getting processor for language:", lang)
	p, ok := f.processors[lang]
	if !ok {
		logger.Logger.Println("Unsupported language requested:", lang)
		return nil, fmt.Errorf("unsupported language: %s", lang)
	}
	logger.Logger.Println("Processor found for language:", lang)
	return p, nil
}
