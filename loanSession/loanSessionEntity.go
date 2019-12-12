package loanSession

import "time"

import "github.com/dhpsagala/loan_process_donnysagala/storage"

import "fmt"

import "strconv"

type ILoanSessionEntity interface {
	CountLoan() int
	GetID() int
	GetDate() time.Time
	GetRequestLimit() int
	GenerateNewLoanID() (int, error)
}

type loanSessionEntity struct {
	ID           int
	Date         time.Time
	RequestLimit int
}

func (ls loanSessionEntity) CountLoan() int {
	var count int
	storage.Db.QueryRow("SELECT COUNT(ID) AS `Count` FROM Loan WHERE SessionID = ?", ls.ID).Scan(&count)
	return count
}

func (ls loanSessionEntity) GetID() int {
	return ls.ID
}

func (ls loanSessionEntity) GetDate() time.Time {
	return ls.Date
}

func (ls loanSessionEntity) GetRequestLimit() int {
	return ls.RequestLimit
}

func (ls loanSessionEntity) GenerateNewLoanID() (int, error) {
	idStr := fmt.Sprintf("%s%d", ls.GetDate().Format("02012006"), (ls.CountLoan() + 1))
	return strconv.Atoi(idStr)
}
