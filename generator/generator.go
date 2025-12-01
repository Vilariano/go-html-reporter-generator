package generator

import (
	"embed"
	"html/template"
	"os"

	"github.com/Vilariano/go-html-reporter-generator/models"
	"github.com/Vilariano/go-html-reporter-generator/utils"
)

//go:embed templates/report.html.tmpl assets/style.css assets/script.js
var reportFS embed.FS

type ReportData struct {
	Features       []models.Feature
	TotalFeatures  int
	TotalScenarios int
	TotalSteps     int
	Passed         int
	Failed         int
	Skipped        int
	// Assets embutidos
	Style  string
	Script string
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
					if scenarioStatus != "failed" {
						scenarioStatus = "skipped"
					}
				}
			}
			features[fi].Elements[ei].Status = scenarioStatus
		}
	}

	// carregar assets embutidos
	styleBytes, _ := reportFS.ReadFile("assets/style.css")
	scriptBytes, _ := reportFS.ReadFile("assets/script.js")
	data.Style = string(styleBytes)
	data.Script = string(scriptBytes)

	// carregar template embutido com funções extras
	tmpl, err := template.New("report.html.tmpl").Funcs(template.FuncMap{
		"nsToMs":      utils.NsToMs,
		"avgDuration": utils.AvgDuration,
		"safeCSS": func(s string) template.CSS {
			return template.CSS(s)
		},
		"safeJS": func(s string) template.JS {
			return template.JS(s)
		},
	}).ParseFS(reportFS, "templates/report.html.tmpl")
	if err != nil {
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	data.Features = features

	return tmpl.ExecuteTemplate(f, "report.html.tmpl", data)
}
