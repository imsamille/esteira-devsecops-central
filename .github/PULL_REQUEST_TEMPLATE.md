## Descrição

<!-- Descreva de forma objetiva o que esta PR entrega e qual PBI/tarefa ela resolve. -->

**PBI/Tarefa:** #
**Tipo de alteração:**
- [ ] Nova funcionalidade
- [ ] Correção de bug
- [ ] Refatoração
- [ ] Documentação
- [ ] Configuração de pipeline/segurança

---

## Contexto de Segurança

<!-- Descreva se a alteração impacta algum componente de segurança da esteira. -->

- [ ] Altera regras do Semgrep (`security/semgrep/rules.yaml`)
- [ ] Altera configuração do Gitleaks (`security/gitleaks/.gitleaks.toml`)
- [ ] Altera steps da pipeline (`pipeline.yml`)
- [ ] Altera scripts de importação para DefectDojo
- [ ] Altera Quality Gate ou Shadow Mode
- [ ] Nenhum impacto em segurança

---

## Evidências

<!-- Cole prints, logs ou links comprovando que a alteração funciona. -->

---

## Checklist

- [ ] Testei localmente antes de abrir a PR
- [ ] Não há credenciais ou secrets no código
- [ ] A documentação foi atualizada (se aplicável)
- [ ] O Shadow Mode permanece intacto (scanners não bloqueiam o build)
- [ ] As regras Semgrep/Gitleaks passam na validação de sintaxe
