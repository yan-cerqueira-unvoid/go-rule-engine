package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yan-cerqueira-unvoid/rules-engine-poc/internal/model"
	"github.com/yan-cerqueira-unvoid/rules-engine-poc/internal/report"
	"github.com/yan-cerqueira-unvoid/rules-engine-poc/internal/rules"
)

func main() {
	age := flag.Int("age", 25, "User age")
	income := flag.Float64("income", 3000.0, "User monthly income")
	creditScore := flag.Int("score", 650, "User credit score")
	hasDebt := flag.Bool("debt", true, "Whether the user has existing debt")
	flag.Parse()

	log.Println("Running credit analysis with verbose logging...")

	userContext := model.NewUserContext(*age, *income, *creditScore, *hasDebt)
	
	engine := rules.NewRuleEngine()
	err := engine.ExecuteRules(userContext)
	if err != nil {
		log.Fatalf("Error executing credit analysis rules: %v", err)
	}

	reporter := report.NewReportGenerator()
	reportOutput := reporter.GenerateReport(userContext)
	fmt.Println(reportOutput)
}
