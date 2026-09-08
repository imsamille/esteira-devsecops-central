# Central DevSecOps - Reusable Workflow (SaaS Interno)

[![Pipeline Base - GitHub Actions](https://github.com/imsamille/esteira-devsecops-central-consolidada/actions/workflows/pipeline.yml/badge.svg)](https://github.com/imsamille/esteira-devsecops-central-consolidada/actions/workflows/pipeline.yml)

Este repositório centraliza uma esteira DevSecOps modular e reutilizável baseada em GitHub Actions e ferramentas open source. Qualquer squad pode plugar seus repositórios à esteira chamando o workflow central, sem duplicar scripts ou configurações complexas.

---

## Como Integrar a Esteira no Seu Projeto

No repositório da sua squad, crie o arquivo `.github/workflows/devsecops-ci.yml` com a estrutura abaixo:

```yaml
name: DevSecOps CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]
  workflow_dispatch:

jobs:
  security:
    uses: imsamille/esteira-devsecops-central-consolidada/.github/workflows/pipeline.yml@main
    with:
      project_name: 'nome-do-seu-projeto'
      stack_type: 'node' # Opções: node, dotnet, python, java, cpp, go, rust
      run_dast: false    # Ative (true) apenas se tiver ambiente web acessível configurado
      dast_target_url: '' # URL alvo para análise dinâmica (obrigatório se run_dast: true)
    secrets:
      DEFECTDOJO_URL: ${{ secrets.DEFECTDOJO_URL }}
      DEFECTDOJO_API_KEY: ${{ secrets.DEFECTDOJO_API_KEY }}
```

### Tabela de Parâmetros

| **Campo**         | **Tipo** | **Obrigatório** | **Padrão** | **Descrição**                                                               |
| ----------------- | -------- | --------------- | ---------- | --------------------------------------------------------------------------- |
| `project_name`    | String   | **Sim**         | -          | Identificador do projeto nos relatórios e no DefectDojo.                    |
| `stack_type`      | String   | **Sim**         | -          | Stack do projeto (`node`, `dotnet`, `python`, `java`, `cpp`, `go`, `rust`). |
| `run_dast`        | Boolean  | Não             | `false`    | Execução de análise dinâmica via OWASP ZAP.                                 |
| `dast_target_url` | String   | Não             | `''`       | URL pública ou container em execução para teste DAST.<br>                   |

### Configuração de Secrets (`secrets`)

Os segredos são opcionais na chamada da esteira:

* **Com DefectDojo integrado:** Cadastre `DEFECTDOJO_URL` e `DEFECTDOJO_API_KEY` na aba **Settings > Secrets and variables > Actions** do repositório da sua aplicação.

* **Sem DefectDojo configurado:** A esteira executa normalmente sem interromper o fluxo. Para evitar erros sintáticos do GitHub Actions, forneça valores dummy diretamente nas Secrets do repositório (`http://placeholder.local` e `dummy-token`), ou declare os secrets vazios na chamada.

## Camadas de Segurança e Ferramentas

=

```text
[Código da Squad]
       │
       ▼
 1. GitLeaks ─────────► Secret Scanning (Detecta tokens, senhas e chaves privadas)
       │
       ▼
 2. Semgrep ──────────► SAST (Análise estática de vulnerabilidades e Dockerfiles)
       │
       ▼
 3. Trivy FS ─────────► SCA (CVEs em bibliotecas terceiras: npm, nuget, pip, etc.)
       │
       ▼
 4. OWASP ZAP ────────► DAST (Análise dinâmica web - ativado sob demanda)
       │
       ▼
[Upload de Relatórios (devsecops-reports.zip)]

```

* **GitLeaks:** Varre histórico e árvore de commits em busca de credenciais em texto puro.
* **Semgrep (SAST Otimizado):** Avalia regras OWASP Top 10 e padrões inseguros com filtros contextuais para a stack indicada, eliminando ruído sintático e arquivos irrelevantes.
* **Trivy (SCA):** Detecta CVEs ativas em lockfiles de dependências (`package-lock.json`, `yarn.lock`, `.csproj`, etc.).
* **OWASP ZAP (DAST):** Scan dinâmico de aplicações web em execução.
* **DefectDojo & Grafana:** Gestão centralizada de vulnerabilidades e dashboards consolidados.

## O que é o Shadow Mode?

Esta esteira roda por padrão em **Shadow Mode** (`exit-code: 0` e tratamento de erros permissivo):

* **Não quebra a esteira produtiva:** O time de desenvolvimento continua trabalhando sem bloqueios imediatos.

* **Gera visibilidade total:** Relatórios JSON completos de cada scanner são empacotados e disponibilizados na aba **Actions > Artifacts** (`devsecops-reports.zip`).

* **Transição para Quality Gate Estrito:** Assim que a squad remediar as vulnerabilidades detectadas, basta remover as flags de tolerância (`|| true`) para converter a esteira em um portão bloqueador de deploy.

## Como Consultar os Relatórios de Vulnerabilidade

1. Acesse a aba **Actions** no repositório da sua aplicação.

2. Clique na execução da pipeline mais recente.

3. No rodapé da página (seção **Artifacts**), baixe o arquivo **`devsecops-reports.zip`**.

4. Descompacte para analisar:

   * `reports/gitleaks/gitleaks-results.json`
   * `reports/semgrep/semgrep-results.json`
   * `reports/trivy/trivy-fs-results.json`
   * `reports/zap/` (quando DAST estiver ativo)

## Equipe Responsável

* Alexia Josielly Duarte da Silva Alves
* João Henrique Lopes de Araújo Freire
* Pedro Henrique Borges Silva
* Raphaela Samille Ramalho de Oliveira
* Thiago Farias Leal
