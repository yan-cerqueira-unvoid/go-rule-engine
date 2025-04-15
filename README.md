# Sistema de Análise de Crédito

Este sistema realiza análise automatizada de crédito usando o motor de regras Grule Rule Engine. O sistema avalia os dados do usuário como idade, renda, score de crédito e existência de dívidas, aplicando regras de negócio para aprovar ou rejeitar solicitações de crédito.

## Principais Funcionalidades

- Análise automática baseada em regras de negócio
- Determinação de limite de crédito com base na renda
- Geração de relatórios detalhados
- Interface de linha de comando com parâmetros configuráveis

## Como Usar

### Compilação

```bash
go build -o creditanalysis ./cmd
```

### Execução

```bash
# Com valores padrão
./creditanalysis

# Especificando parâmetros
./creditanalysis -age=30 -income=4500 -score=720 -debt=false -verbose
```

### Parâmetros Disponíveis

- `-age`: Idade do solicitante (padrão: 25)
- `-income`: Renda mensal (padrão: 3000.0)
- `-score`: Score de crédito (padrão: 650)
- `-debt`: Possui dívidas (padrão: true)
- `-verbose`: Exibe logs detalhados (padrão: false)

## Estrutura do Projeto

- `cmd/app/main.go`: Ponto de entrada da aplicação
- `internal/domain/user.go`: Definição do contexto do usuário
- `internal/rules/`: Motor de regras e definições
- `internal/report/`: Geração de relatórios

## Regras de Negócio

As regras de crédito incluem:
1. Idade mínima de 18 anos
2. Renda mínima de R$ 1.000,00
3. Score de crédito mínimo de 500
4. Com dívidas existentes, score mínimo de 700
5. Limite de crédito baseado na faixa de renda

## Requisitos

- Go 1.18 ou superior
- Dependências gerenciadas via Go Modules


