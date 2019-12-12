package loan

import (
	"time"

	"github.com/dhpsagala/loan_process_donnysagala/storage"
)

func New(loanID int, sessionID int, customerID int, amount int, isApproved bool) (ILoanEntity, error) {
	stmt, err := storage.Db.Prepare(
		`INSERT INTO Loan (LoanID, SessionID, CustomerID, Amount, CreateDate, IsApproved) VALUES (?, ?, ?, ?, DATE('now'), ?)`,
	)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(loanID, sessionID, customerID, amount, isApproved)
	if err != nil {
		return nil, err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return GetByLoanID(loanID)
}

func GetByLoanID(loanID int) (ILoanEntity, error) {
	rows, err := storage.Db.Query("SELECT * FROM Loan WHERE LoanID = ? LIMIT 1", loanID)
	if err != nil {
		return nil, err
	}

	datas := []ILoanEntity{}
	defer rows.Close()
	for rows.Next() {
		var ID, LoanID, SessionID, CustomerID, Amount int
		var CreateDate time.Time
		var IsApproved bool

		err := rows.Scan(&ID, &LoanID, &SessionID, &CustomerID, &Amount, &CreateDate, &IsApproved)
		if err != nil {
			return nil, err
		}

		var p ILoanEntity = loanEntity{ID: ID, LoanID: LoanID, SessionID: SessionID, CustomerID: CustomerID, Amount: Amount, CreateDate: CreateDate, IsApproved: IsApproved}
		datas = append(datas, p)
	}

	if len(datas) > 0 {
		return datas[0], nil
	}
	return nil, nil
}

func GetBySessionIdAndAmountAndStatus(sessionID int, amount int, isApproved bool) ([]ILoanEntity, error) {
	rows, err := storage.Db.Query(
		"SELECT * FROM Loan WHERE SessionID = ? AND Amount = ? AND IsApproved = ?",
		sessionID,
		amount,
		isApproved,
	)
	if err != nil {
		return nil, err
	}

	datas := []ILoanEntity{}
	defer rows.Close()
	for rows.Next() {
		var ID, LoanID, SessionID, CustomerID, Amount int
		var CreateDate time.Time
		var IsApproved bool

		err := rows.Scan(&ID, &LoanID, &SessionID, &CustomerID, &Amount, &CreateDate, &IsApproved)
		if err != nil {
			return nil, err
		}

		var p ILoanEntity = loanEntity{ID: ID, LoanID: LoanID, SessionID: SessionID, CustomerID: CustomerID, Amount: Amount, CreateDate: CreateDate, IsApproved: IsApproved}
		datas = append(datas, p)
	}
	return datas, nil
}
