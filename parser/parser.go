package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Vilariano/go-html-reporter-generator/models"
)

// ParseJSON lê um arquivo JSON no formato Cucumber/Godog e retorna uma lista de Features
func ParseJSON(file string) ([]models.Feature, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo %s: %w", file, err)
	}

	var features []models.Feature
	if err := json.Unmarshal(data, &features); err != nil {
		return nil, fmt.Errorf("erro ao interpretar JSON: %w", err)
	}

	// Garante que nunca retorna nil
	if features == nil {
		features = []models.Feature{}
	}

	return features, nil
}
