package parser

import (
	"encoding/json"
	"os"

	"github.com/Vilariano/go-html-reporter-generator/internal/models"
)

func ParseJSON(file string) ([]models.Feature, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var features []models.Feature
	if err := json.Unmarshal(data, &features); err != nil {
		return nil, err
	}

	return features, nil
}
