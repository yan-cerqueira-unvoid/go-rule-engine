package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// UserContext contains the data to be analyzed by the rules
type UserContext struct {
	// Input data
	Age         int
	Income      float64
	CreditScore int
	HasDebt     bool

	// Analysis results
	CreditApproved   bool
	CreditRejected   bool
	CreditLimit      float64
	RejectionReasons []string
}

func (uc *UserContext) AddRejectionReason(reason string) {
	if !slices.Contains(uc.RejectionReasons, reason) {
		uc.RejectionReasons = append(uc.RejectionReasons, reason)
	}
}

func (uc *UserContext) PrintReport() {
	fmt.Println("=== CREDIT ANALYSIS REPORT ===")
	fmt.Printf("Age: %d\n", uc.Age)
	fmt.Printf("Income: $ %.2f\n", uc.Income)
	fmt.Printf("Credit Score: %d\n", uc.CreditScore)
	fmt.Printf("Has Debt: %t\n\n", uc.HasDebt)

	if uc.CreditApproved {
		fmt.Println("RESULT: APPROVED ✓")
		fmt.Printf("Credit Limit: $ %.2f\n", uc.CreditLimit)
	} else {
		fmt.Println("RESULT: REJECTED ✗")
		fmt.Println("Reasons:")
		for _, reason := range uc.RejectionReasons {
			fmt.Printf("  - %s\n", reason)
		}
	}
}

func main() {
	userContext := &UserContext{
		Age:              25,
		Income:           3000.0,
		CreditScore:      650,
		HasDebt:          true,
		RejectionReasons: []string{}, // Initialize with empty slice
	}

	// Initialize the rules engine
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	// Define the rules
	drl := `
rule AgeCheck "Checks if the user has the minimum age" salience 10 {
    when
        UserContext.Age < 18
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Minimum age to request credit is 18 years");
}

rule LowIncomeCheck "Checks if income is sufficient" salience 9 {
    when
        UserContext.Income < 1000 && !UserContext.CreditRejected
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Insufficient income (minimum of $1,000.00)");
}

rule BadCreditScoreCheck "Checks if credit score is acceptable" salience 8 {
    when
        UserContext.CreditScore < 500 && !UserContext.CreditRejected
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Credit score too low (minimum 500)");
}

rule DebtAndLowScoreCheck "Checks combination of debt and low score" salience 7 {
    when
        UserContext.HasDebt == true && 
        UserContext.CreditScore < 700 && 
        !UserContext.CreditRejected
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Combination of existing debt and credit score less than 700");
}

rule CreditLimitHighIncome "Sets credit limit for high income" salience 5 {
    when
        UserContext.Income >= 5000 && 
        !UserContext.CreditRejected && 
        UserContext.CreditLimit == 0
    then
        UserContext.CreditLimit = UserContext.Income * 2;
}

rule CreditLimitMediumIncome "Sets credit limit for medium income" salience 4 {
    when
        UserContext.Income >= 2000 && 
        UserContext.Income < 5000 && 
        !UserContext.CreditRejected && 
        UserContext.CreditLimit == 0
    then
        UserContext.CreditLimit = UserContext.Income * 1.5;
}

rule CreditLimitLowIncome "Sets credit limit for low income" salience 3 {
    when
        UserContext.Income >= 1000 && 
        UserContext.Income < 2000 && 
        !UserContext.CreditRejected && 
        UserContext.CreditLimit == 0
    then
        UserContext.CreditLimit = UserContext.Income;
}

rule ApproveCredit "Approves credit if not rejected" salience 0 {
    when
        !UserContext.CreditRejected && UserContext.CreditLimit > 0
    then
        UserContext.CreditApproved = true;
}
`

	// Add the rules to the engine
	err := ruleBuilder.BuildRuleFromResource("CreditRules", "1.0.0", pkg.NewBytesResource([]byte(drl)))
	if err != nil {
		log.Fatalf("Error building rules: %v", err)
	}

	// Get the knowledge base
	knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance("CreditRules", "1.0.0")
	if err != nil {
		log.Fatalf("Error retrieving knowledge base: %v", err)
	}

	// Create the data context
	dataContext := ast.NewDataContext()
	err = dataContext.Add("UserContext", userContext)
	if err != nil {
		log.Fatalf("Error adding user context: %v", err)
	}

	// Create and execute the rules engine
	ruleEngine := engine.NewGruleEngine()
	ruleEngine.MaxCycle = 100 // Set a limit to avoid infinite loops

	err = ruleEngine.Execute(dataContext, knowledgeBase)
	if err != nil {
		log.Fatalf("Error executing rules: %v", err)
	}

	// Print the report
	userContext.PrintReport()
}
