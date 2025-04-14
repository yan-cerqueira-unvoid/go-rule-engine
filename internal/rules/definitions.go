package rules

func GetRuleDefinitions() string {
	return `
rule AgeCheck "Verifica se o usuário tem a idade mínima" salience 10 {
    when
        UserContext.Age < 18
    then
        UserContext.SetRejected(true);
        UserContext.AddRejectionReason("Idade mínima para solicitar crédito é 18 anos");
}

rule LowIncomeCheck "Verifica se a renda é suficiente" salience 9 {
    when
        UserContext.Income < 1000 && !UserContext.IsRejected()
    then
        UserContext.SetRejected(true);
        UserContext.AddRejectionReason("Renda insuficiente (mínimo de R$ 1.000,00)");
}

rule BadCreditScoreCheck "Verifica se o score de crédito é aceitável" salience 8 {
    when
        UserContext.CreditScore < 500 && !UserContext.IsRejected()
    then
        UserContext.SetRejected(true);
        UserContext.AddRejectionReason("Score de crédito muito baixo (mínimo 500)");
}

rule DebtAndLowScoreCheck "Verifica combinação de dívida e score baixo" salience 7 {
    when
        UserContext.HasDebt == true && 
        UserContext.CreditScore < 700 && 
        !UserContext.IsRejected()
    then
        UserContext.SetRejected(true);
        UserContext.AddRejectionReason("Combinação de dívida existente e score de crédito menor que 700");
}

rule CreditLimitHighIncome "Define limite de crédito para alta renda" salience 5 {
    when
        UserContext.Income >= 5000 && 
        !UserContext.IsRejected() && 
        UserContext.GetCreditLimit() == 0
    then
        UserContext.SetCreditLimit(UserContext.Income * 2);
}

rule CreditLimitMediumIncome "Define limite de crédito para renda média" salience 4 {
    when
        UserContext.Income >= 2000 && 
        UserContext.Income < 5000 && 
        !UserContext.IsRejected() && 
        UserContext.GetCreditLimit() == 0
    then
        UserContext.SetCreditLimit(UserContext.Income * 1.5);
}

rule CreditLimitLowIncome "Define limite de crédito para baixa renda" salience 3 {
    when
        UserContext.Income >= 1000 && 
        UserContext.Income < 2000 && 
        !UserContext.IsRejected() && 
        UserContext.GetCreditLimit() == 0
    then
        UserContext.SetCreditLimit(UserContext.Income);
}

rule ApproveCredit "Aprova crédito se não rejeitado" salience 0 {
    when
        !UserContext.IsRejected() && UserContext.GetCreditLimit() > 0
    then
        UserContext.SetApproved(true);
}
`
}
