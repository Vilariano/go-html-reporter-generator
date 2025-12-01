package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Vilariano/go-html-reporter-generator/parser"
	"github.com/Vilariano/go-html-reporter-generator/reporter"
)

func main() {
	// Flags para entrada e saída
	input := flag.String("input", "results.json", "Arquivo JSON de entrada")
	output := flag.String("output", "report.html", "Arquivo HTML de saída")
	flag.Parse()

	// Parse do arquivo JSON
	features, err := parser.ParseJSON(*input)
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		os.Exit(1)
	}

	// Geração do relatório via API pública
	if err := reporter.GenerateReport(features, *output); err != nil {
		fmt.Println("Erro ao gerar HTML:", err)
		os.Exit(1)
	}

	fmt.Println("Relatório gerado em:", *output)
}
