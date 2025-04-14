package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/yan-cerqueira-unvoid/rules-engine-poc/internal/model"
)

type ReportGenerator struct {
	timeFormat string
}

func NewReportGenerator() *ReportGenerator {
	return &ReportGenerator{
		timeFormat: "02/01/2006 15:04:05", // Formato de data no padrão brasileiro
	}
}

func (rg *ReportGenerator) GenerateReport(userContext *model.UserContext) string {
	var b strings.Builder

	now := time.Now()
	b.WriteString("══════════════════════════════════════════\n")
	b.WriteString("       RELATÓRIO DE ANÁLISE DE CRÉDITO    \n")
	b.WriteString("══════════════════════════════════════════\n")
	b.WriteString(fmt.Sprintf("Data: %s\n\n", now.Format(rg.timeFormat)))

	b.WriteString("INFORMAÇÕES DO SOLICITANTE\n")
	b.WriteString("──────────────────────────\n")
	b.WriteString(fmt.Sprintf("Idade: %d anos\n", userContext.Age))
	b.WriteString(fmt.Sprintf("Renda mensal: R$ %.2f\n", userContext.Income))
	b.WriteString(fmt.Sprintf("Score de crédito: %d\n", userContext.CreditScore))
	
	if userContext.HasDebt {
		b.WriteString("Possui dívidas: Sim\n\n")
	} else {
		b.WriteString("Possui dívidas: Não\n\n")
	}

	b.WriteString("RESULTADO DA ANÁLISE\n")
	b.WriteString("───────────────────\n")
	
	if userContext.IsApproved() {
		b.WriteString("CRÉDITO APROVADO ✓\n")
		b.WriteString(fmt.Sprintf("Limite de crédito: R$ %.2f\n", userContext.GetCreditLimit()))
		
		b.WriteString("\nSUGESTÕES DE PRODUTOS\n")
		b.WriteString("────────────────────\n")
		
		if userContext.GetCreditLimit() > 5000 {
			b.WriteString("✓ Cartão Platinum\n")
			b.WriteString("✓ Empréstimo pessoal\n")
			b.WriteString("✓ Financiamento imobiliário\n")
		} else if userContext.GetCreditLimit() > 2000 {
			b.WriteString("✓ Cartão Gold\n")
			b.WriteString("✓ Empréstimo pessoal\n")
		} else {
			b.WriteString("✓ Cartão Standard\n")
		}
	} else {
		b.WriteString("CRÉDITO REJEITADO ✗\n\n")
		b.WriteString("Motivos da rejeição:\n")
		
		for i, reason := range userContext.GetRejectionReasons() {
			b.WriteString(fmt.Sprintf("%d. %s\n", i+1, reason))
		}
		
		b.WriteString("\nRECOMENDAÇÕES\n")
		b.WriteString("─────────────\n")
		b.WriteString("• Verifique se há pendências financeiras\n")
		b.WriteString("• Considere aumentar sua renda\n")
		b.WriteString("• Trabalhe para melhorar seu score de crédito\n")
	}

	b.WriteString("\n══════════════════════════════════════════\n")
	b.WriteString("Análise realizada por Sistema Automatizado\n")
	b.WriteString("Este relatório é apenas informativo\n")
	
	return b.String()
}
