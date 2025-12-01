// Script para interatividade do relatório

console.log("Relatório carregado com sucesso!");

// Expandir/Recolher todos os cenários
function toggleAllScenarios(expand = true) {
  const scenarios = document.querySelectorAll("details.scenario");
  scenarios.forEach(s => {
    s.open = expand;
  });
}

// Filtrar cenários por status
function filterScenarios(status) {
  const rows = document.querySelectorAll("tbody tr");
  rows.forEach(row => {
    const statusCell = row.querySelector("td:nth-child(2)");
    if (!status || statusCell.classList.contains(status)) {
      row.style.display = "";
    } else {
      row.style.display = "none";
    }
  });
}

// Adicionar botões extras na navbar
document.addEventListener("DOMContentLoaded", () => {
  const navbar = document.querySelector("nav.navbar div");
  if (navbar) {
    const btnExpand = document.createElement("button");
    btnExpand.textContent = "Expandir Todos";
    btnExpand.onclick = () => toggleAllScenarios(true);

    const btnCollapse = document.createElement("button");
    btnCollapse.textContent = "Recolher Todos";
    btnCollapse.onclick = () => toggleAllScenarios(false);

    const btnShowPassed = document.createElement("button");
    btnShowPassed.textContent = "Mostrar Passados";
    btnShowPassed.onclick = () => filterScenarios("passed");

    const btnShowFailed = document.createElement("button");
    btnShowFailed.textContent = "Mostrar Falhados";
    btnShowFailed.onclick = () => filterScenarios("failed");

    const btnShowAll = document.createElement("button");
    btnShowAll.textContent = "Mostrar Todos";
    btnShowAll.onclick = () => filterScenarios(null);

    navbar.appendChild(btnExpand);
    navbar.appendChild(btnCollapse);
    navbar.appendChild(btnShowPassed);
    navbar.appendChild(btnShowFailed);
    navbar.appendChild(btnShowAll);
  }
});
