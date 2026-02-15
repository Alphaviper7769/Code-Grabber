package domain

type ProcessorFactory interface {
	GetProcessor(language string) (Processor, error)
}
