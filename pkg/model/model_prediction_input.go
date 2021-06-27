package model

type PredictionInput struct {
	Champions   []string `json:"champions"`
	Intercept   bool     `json:"intercept"`
	Shuffle     bool     `json:"shuffle"`
	Serendipity bool     `json:"serendipity"`
}

