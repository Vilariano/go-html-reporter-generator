# go-html-reporter-generator

Gerador de relatórios em **HTML** para resultados de testes automatizados (ex.: Cucumber JSON).  
O objetivo é transformar arquivos de saída (`results.json`) em relatórios visuais e interativos, com métricas, gráficos e detalhes de execução.

---

## 🚀 Funcionalidades

- 📊 **Dashboard de métricas**: total de features, cenários e passos, além de status (passaram, falharam, ignorados).
- 📈 **Gráficos interativos** (Chart.js):
  - Pizza com distribuição de status.
  - Barras com tempo médio por cenário.
- 📝 **Detalhamento por Feature e Cenário**:
  - Lista de passos com status, duração e localização.
  - Badge de status por cenário.
- 🎨 **Layout moderno e responsivo**:
  - Navbar fixa com exportação para PDF.
  - Sidebar com navegação rápida entre features.
  - Cards coloridos para métricas e cenários.

---

## 📦 Instalação
- Clone o repositório:
    ```bash
    git clone https://github.com/seu-usuario/go-html-reporter-generator.git
    cd go-html-reporter-generator
    ```

- Instale as dependências:
    ```bash
    go mod tidy
    ```

## ▶️ Uso
- Execute o gerador passando o arquivo JSON de resultados e o nome do HTML de saída:
    ```bash
    go run cmd/reporter/main.go --input results.json --output report.html
    ```
    Isso irá gerar um arquivo report.html pronto para abrir no navegador.


## 📂 Estrutura do projeto
```Código
go-html-reporter-generator/
├── cmd/
│   └── reporter/
│       └── main.go         # Ponto de entrada do CLI
├── internal/
│   ├── generator/          # Lógica de geração do HTML
│   │   └── generator.go
│   ├── models/             # Estruturas de dados (Feature, Element, Step)
│   │   └── models.go
│   └── utils/              # Funções auxiliares (conversão, cálculos)
│       └── utils.go
├── templates/
│   └── report.html.tmpl    # Template HTML do relatório
├── assets/
│   ├── style.css           # Estilos customizados
│   └── script.js           # Scripts adicionais
└── README.md
```

## 🛠️ Desenvolvimento
- Adicionar novas métricas
    - Alterar ReportData em generator.go.
    - Atualizar o template report.html.tmpl.

- Customizar layout
    Editar assets/style.css para cores, fontes e posicionamento.
    Ajustar report.html.tmpl para novos componentes.

- Funções auxiliares
    - utils.NsToMs: converte nanosegundos em milissegundos.
    - utils.AvgDuration: calcula tempo médio dos passos de um cenário.

## 📊 Exemplo de relatório
- Cards de métricas no topo.
- Gráficos lado a lado.
- Features listadas com cenários e passos detalhados.
- Exportação rápida para PDF via botão na navbar.

## 🤝 Contribuição
1. Faça um fork do projeto.
2. Crie uma branch para sua feature (git checkout -b minha-feature).
3. Commit suas alterações (git commit -m 'Adicionei nova feature').
4. Push para a branch (git push origin minha-feature).
5. Abra um Pull Request.
