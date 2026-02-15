package factory

import (
	"fmt"
	"native/internal/domain"
	"native/internal/infrastructure/processors"
)

type LanguageProcessorFactory struct {
	processors map[string]domain.Processor
}

func NewLanguageProcessorFactory(baseDir string) *LanguageProcessorFactory {

	return &LanguageProcessorFactory{
		processors: map[string]domain.Processor{
			"cpp": processors.NewCppProcessor(baseDir),
		},
	}
}

func (f *LanguageProcessorFactory) GetProcessor(lang string) (domain.Processor, error) {

	p, ok := f.processors[lang]
	if !ok {
		return nil, fmt.Errorf("unsupported language: %s", lang)
	}

	return p, nil
}
