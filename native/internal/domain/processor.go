package domain

type Processor interface {
	Process(problem Problem) error
}
