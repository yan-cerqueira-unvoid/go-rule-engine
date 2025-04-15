package rules

// GetRuleDefinitions retorna as definições de regras no formato DRL
func GetRuleDefinitions() string {
	return `
// Importante: Voltamos a usar os campos diretamente em vez dos métodos getter/setter
// já que o Grule pode ter problemas com eles em versões específicas

rule AgeCheck "Verifica se o usuário tem a idade mínima" salience 10 {
    when
        UserContext.Age < 18
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Idade mínima para solicitar crédito é 18 anos");
        Retract("AgeCheck"); // Remove esta regra do pool de execução após executá-la
}

rule LowIncomeCheck "Verifica se a renda é suficiente" salience 9 {
    when
        UserContext.Income < 1000 && UserContext.CreditRejected == false
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Renda insuficiente (mínimo de R$ 1.000,00)");
        Retract("LowIncomeCheck"); // Remove esta regra após executá-la
}

rule BadCreditScoreCheck "Verifica se o score de crédito é aceitável" salience 8 {
    when
        UserContext.CreditScore < 500 && UserContext.CreditRejected == false
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Score de crédito muito baixo (mínimo 500)");
        Retract("BadCreditScoreCheck"); // Remove esta regra após executá-la
}

rule DebtAndLowScoreCheck "Verifica combinação de dívida e score baixo" salience 7 {
    when
        UserContext.HasDebt == true && 
        UserContext.CreditScore < 700 && 
        UserContext.CreditRejected == false
    then
        UserContext.CreditRejected = true;
        UserContext.AddRejectionReason("Combinação de dívida existente e score de crédito menor que 700");
        Retract("DebtAndLowScoreCheck"); // Remove esta regra após executá-la
}

rule CreditLimitHighIncome "Define limite de crédito para alta renda" salience 5 {
    when
        UserContext.Income >= 5000 && 
        UserContext.CreditRejected == false && 
        UserContext.CreditLimit == 0
    then
        UserContext.CreditLimit = UserContext.Income * 2;
        Retract("CreditLimitHighIncome"); // Remove esta regra após executá-la
}

rule CreditLimitMediumIncome "Define limite de crédito para renda média" salience 4 {
    when
        UserContext.Income >= 2000 && 
        UserContext.Income < 5000 && 
        UserContext.CreditRejected == false && 
        UserContext.CreditLimit == 0
    then
        UserContext.CreditLimit = UserContext.Income * 1.5;
        Retract("CreditLimitMediumIncome"); // Remove esta regra após executá-la
}

rule CreditLimitLowIncome "Define limite de crédito para baixa renda" salience 3 {
    when
        UserContext.Income >= 1000 && 
        UserContext.Income < 2000 && 
        UserContext.CreditRejected == false && 
        UserContext.CreditLimit == 0
    then
        UserContext.CreditLimit = UserContext.Income;
        Retract("CreditLimitLowIncome"); // Remove esta regra após executá-la
}

rule ApproveCredit "Aprova crédito se não rejeitado" salience 0 {
    when
        UserContext.CreditRejected == false && UserContext.CreditLimit > 0
    then
        UserContext.CreditApproved = true;
        Retract("ApproveCredit"); // Remove esta regra após executá-la
}`
}
