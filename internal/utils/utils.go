package utils

import (
	"github.com/Vilariano/go-html-reporter-generator/internal/models"
)

// Converte nanosegundos para milissegundos
func NsToMs(ns int64) float64 {
	return float64(ns) / 1_000_000.0
}

// Calcula a média de duração dos passos de um cenário
func AvgDuration(steps []models.Step) float64 {
	if len(steps) == 0 {
		return 0
	}
	var total int64
	for _, s := range steps {
		total += s.Result.Duration
	}
	return float64(total) / float64(len(steps)) / 1_000_000.0
}
