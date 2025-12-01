package reporter

import (
	"github.com/Vilariano/go-html-reporter-generator/generator"
	"github.com/Vilariano/go-html-reporter-generator/models"
)

func GenerateReport(features []models.Feature, output string) error {
	return generator.GenerateHTML(features, output)
}
