package loan

import "time"

import "fmt"

type ILoanEntity interface {
	GetID() int
	GetLoanID() int
	GetSessionID() int
	GetAmount() int
	StatusText() string
}

type loanEntity struct {
	ID         int
	LoanID     int
	SessionID  int
	CustomerID int
	Amount     int
	CreateDate time.Time
	IsApproved bool
}

func (l loanEntity) GetID() int {
	return l.ID
}

func (l loanEntity) GetLoanID() int {
	return l.LoanID
}

func (l loanEntity) GetSessionID() int {
	return l.SessionID
}

func (l loanEntity) GetAmount() int {
	return l.Amount
}

func (l loanEntity) StatusText() string {
	if l.IsApproved {
		return fmt.Sprintf("Loan ID %d is Accepted", l.ID)
	}
	return fmt.Sprintf("Loan ID %d is Rejected", l.ID)
}
