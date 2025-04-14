package rules

import (
	"fmt"
	"log"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/yan-cerqueira-unvoid/rules-engine-poc/internal/model"
)

const (
	RulesetName    = "CreditRules"
	RulesetVersion = "1.0.0"
	MaxCycles      = 100 // Limit to avoid infinity loop
)

type RuleEngine struct {
	knowledgeLibrary *ast.KnowledgeLibrary
	ruleBuilder      *builder.RuleBuilder
	gruleEngine      *engine.GruleEngine
}

func NewRuleEngine() *RuleEngine {
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
	gruleEngine := engine.NewGruleEngine()
	gruleEngine.MaxCycle = MaxCycles

	return &RuleEngine{
		knowledgeLibrary: knowledgeLibrary,
		ruleBuilder:      ruleBuilder,
		gruleEngine:      gruleEngine,
	}
}

func (re *RuleEngine) ExecuteRules(userContext *model.UserContext) error {
	err := re.loadRules()
	if err != nil {
		return fmt.Errorf("failed to load rules: %w", err)
	}

	knowledgeBase, err := re.knowledgeLibrary.NewKnowledgeBaseInstance(RulesetName, RulesetVersion)
	if err != nil {
		return fmt.Errorf("failed to create knowledge base: %w", err)
	}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("UserContext", userContext)
	if err != nil {
		return fmt.Errorf("failed to add user context to data context: %w", err)
	}

	log.Println("Executing credit analysis rules...")
	
    err = re.gruleEngine.Execute(dataContext, knowledgeBase)
	if err != nil {
		return fmt.Errorf("rule execution failed: %w", err)
	}
	
    log.Println("Rules execution completed successfully")

	return nil
}

func (re *RuleEngine) loadRules() error {
	drl := GetRuleDefinitions()
	
	err := re.ruleBuilder.BuildRuleFromResource(RulesetName, RulesetVersion, 
		pkg.NewBytesResource([]byte(drl)))
	if err != nil {
		return fmt.Errorf("failed to build rules from resource: %w", err)
	}
	
	return nil
}
