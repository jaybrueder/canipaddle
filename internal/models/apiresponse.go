package models

type APIResponse struct {
	Result   string    `json:"result"`
	Measures []Measure `json:"measures"`
}
