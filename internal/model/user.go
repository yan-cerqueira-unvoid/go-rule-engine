package domain

import (
	"slices"
	"sync"
)

type UserContext struct {
	mu sync.RWMutex // Protected concurent access

	// Input
	Age         int     `json:"age"`
	Income      float64 `json:"income"`
	CreditScore int     `json:"creditScore"`
	HasDebt     bool    `json:"hasDebt"`

	// Results
	CreditApproved   bool     `json:"creditApproved"`
	CreditRejected   bool     `json:"creditRejected"`
	CreditLimit      float64  `json:"creditLimit"`
	RejectionReasons []string `json:"rejectionReasons"`
}

func NewUserContext(age int, income float64, creditScore int, hasDebt bool) *UserContext {
	return &UserContext{
		Age:              age,
		Income:           income,
		CreditScore:      creditScore,
		HasDebt:          hasDebt,
		RejectionReasons: make([]string, 0),
	}
}

func (uc *UserContext) AddRejectionReason(reason string) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	
	if !slices.Contains(uc.RejectionReasons, reason) {
		uc.RejectionReasons = append(uc.RejectionReasons, reason)
	}
}

func (uc *UserContext) GetRejectionReasons() []string {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	
	result := make([]string, len(uc.RejectionReasons))
	copy(result, uc.RejectionReasons)
	return result
}

func (uc *UserContext) SetCreditLimit(limit float64) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	
	uc.CreditLimit = limit
}

func (uc *UserContext) GetCreditLimit() float64 {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	
	return uc.CreditLimit
}

func (uc *UserContext) IsRejected() bool {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	
	return uc.CreditRejected
}

func (uc *UserContext) SetRejected(rejected bool) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	
	uc.CreditRejected = rejected
}

func (uc *UserContext) SetApproved(approved bool) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	
	uc.CreditApproved = approved
}

func (uc *UserContext) IsApproved() bool {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	
	return uc.CreditApproved
}
