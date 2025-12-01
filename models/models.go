package models

type StepResult struct {
	Status   string `json:"status"`
	Duration int64  `json:"duration"`
}

type StepMatch struct {
	Location string `json:"location"`
}

type Step struct {
	Keyword string     `json:"keyword"`
	Name    string     `json:"name"`
	Line    int        `json:"line"`
	Match   StepMatch  `json:"match"`
	Result  StepResult `json:"result"`
}

type Element struct {
	ID          string `json:"id"`
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Line        int    `json:"line"`
	Type        string `json:"type"`
	Steps       []Step `json:"steps"`
	Status      string `json:"-"` // campo calculado, não vem do JSON
}

type Feature struct {
	URI         string    `json:"uri"`
	ID          string    `json:"id"`
	Keyword     string    `json:"keyword"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Line        int       `json:"line"`
	Elements    []Element `json:"elements"`
}
