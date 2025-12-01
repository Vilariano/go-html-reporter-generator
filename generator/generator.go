package generator

import (
	"html/template"
	"os"

	"github.com/Vilariano/go-html-reporter-generator/models"
	"github.com/Vilariano/go-html-reporter-generator/utils"
)

type ReportData struct {
	Features       []models.Feature
	TotalFeatures  int
	TotalScenarios int
	TotalSteps     int
	Passed         int
	Failed         int
	Skipped        int
}

func GenerateHTML(features []models.Feature, output string) error {
	// calcular métricas
	data := ReportData{
		Features:      features,
		TotalFeatures: len(features),
	}

	for fi, f := range features {
		for ei, e := range f.Elements {
			data.TotalScenarios++

			// calcular status do cenário
			scenarioStatus := "passed"
			for _, s := range e.Steps {
				data.TotalSteps++
				switch s.Result.Status {
				case "passed":
					data.Passed++
				case "failed":
					data.Failed++
					scenarioStatus = "failed"
				case "skipped":
					data.Skipped++
					// só marca como skipped se não houver falha
					if scenarioStatus != "failed" {
						scenarioStatus = "skipped"
					}
				}
			}
			features[fi].Elements[ei].Status = scenarioStatus
		}
	}

	// carregar template com funções extras
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"nsToMs":      utils.NsToMs,
		"avgDuration": utils.AvgDuration,
	}).ParseFiles("templates/report.html.tmpl")
	if err != nil {
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	// usar a versão atualizada de features com Status calculado
	data.Features = features

	return tmpl.ExecuteTemplate(f, "report.html.tmpl", data)
}
