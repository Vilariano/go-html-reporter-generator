# go-html-reporter-generator
## Gerador de relatórios HTML para testes BDD (Cucumber/Godog). Transforma arquivos JSON no formato Cucumber em relatórios HTML interativos e autossuficientes (com CSS e JS embutidos).

### 🚀 Instalação
- Adicione ao seu projeto Go:
  ```go
  go get github.com/Vilariano/go-html-reporter-generator@latest
  ```
  Ou fixe uma versão específica:
  ```go
  go get github.com/Vilariano/go-html-reporter-generator@v1.0.2
  ```

### 📊 Uso via CLI
- Se você quiser rodar direto pela linha de comando:
  ```go
  go run ./cmd/reporter --input report.json --output report.html
  ```
- --input: caminho para o JSON gerado pelo Cucumber/Godog
- --output: caminho do relatório HTML que será criado

### 🧩 Uso como biblioteca
- Você também pode usar no seu código Go:
package main
  ```go
  import (
      "encoding/json"
      "os"

      "github.com/Vilariano/go-html-reporter-generator/models"
      "github.com/Vilariano/go-html-reporter-generator/reporter"
  )

  func main() {
      // Lê o JSON gerado pelo Cucumber/Godog
      file, _ := os.Open("report.json")
      defer file.Close()

      var features []models.Feature
      json.NewDecoder(file).Decode(&features)

      // Gera o relatório HTML
      reporter.GenerateReport(features, "report.html")
  }
  ```

### 🧪 Exemplo com Godog
- Um exemplo de teste integrado (bdd_test.go):
  ```go
  package main

  import (
      "encoding/json"
      "os"
      "testing"

      "github.com/Vilariano/go-html-reporter-generator/models"
      "github.com/Vilariano/go-html-reporter-generator/reporter"
      "github.com/cucumber/godog"
  )

  func TestFeatures(t *testing.T) {
      file, _ := os.Create("report.json")
      defer file.Close()

      opts := godog.Options{
          Format: "cucumber",
          Paths:  []string{"features"},
          Output: file,
      }

      suite := godog.TestSuite{
          Name:                "petstore",
          ScenarioInitializer: InitializeScenario,
          Options:             &opts,
      }

      if suite.Run() != 0 {
          t.Fail()
      }

      jsonFile, _ := os.Open("report.json")
      defer jsonFile.Close()

      var features []models.Feature
      json.NewDecoder(jsonFile).Decode(&features)

      reporter.GenerateReport(features, "report.html")
  }
  ```

### 🎨 Relatório gerado
- Dashboard com métricas (features, cenários, passos, status)
- Gráficos interativos (pizza e barras) usando Chart.js
- Cenários detalhados com passos, status e duração
- Botões para expandir/recolher cenários e filtrar por status
- Exportação para PDF via botão na interface