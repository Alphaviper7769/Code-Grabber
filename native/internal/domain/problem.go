package domain

type Problem struct {
	Slug     string     `json:"slug"`
	Title    string     `json:"title"`
	URL      string     `json:"url"`
	Source   string     `json:"source"`
	Language string     `json:"language"`
	Tests    []TestCase `json:"tests"`
}

type TestCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}
