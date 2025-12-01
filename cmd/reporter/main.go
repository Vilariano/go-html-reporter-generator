package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Vilariano/go-html-reporter-generator/internal/generator"
	"github.com/Vilariano/go-html-reporter-generator/internal/parser"
)

func main() {
	input := flag.String("input", "results.json", "Arquivo JSON de entrada")
	output := flag.String("output", "report.html", "Arquivo HTML de saída")
	flag.Parse()

	features, err := parser.ParseJSON(*input)
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		os.Exit(1)
	}

	if err := generator.GenerateHTML(features, *output); err != nil {
		fmt.Println("Erro ao gerar HTML:", err)
		os.Exit(1)
	}

	fmt.Println("Relatório gerado em:", *output)
}
