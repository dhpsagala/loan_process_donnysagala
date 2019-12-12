package loanSession

import (
	"time"

	"github.com/dhpsagala/loan_process_donnysagala/storage"
)

func Load(limit int) (ILoanSessionEntity, error) {
	curr, err := GetCurrent()
	if err != nil {
		return nil, err
	}

	if curr != nil {
		return curr, nil
	}
	insertSession(limit)
	return GetCurrent()
}

func GetCurrent() (ILoanSessionEntity, error) {
	rows, err := storage.Db.Query("SELECT * FROM LoanSession WHERE Date = ? LIMIT 1", getTodayDate())
	if err != nil {
		return nil, err
	}

	datas := []ILoanSessionEntity{}
	defer rows.Close()
	for rows.Next() {
		var ID, RequestLimit int
		var Date time.Time

		err := rows.Scan(&ID, &Date, &RequestLimit)
		if err != nil {
			return nil, err
		}

		var ls ILoanSessionEntity = loanSessionEntity{ID: ID, Date: Date, RequestLimit: RequestLimit}
		datas = append(datas, ls)
	}

	if len(datas) > 0 {
		return datas[0], nil
	}
	return nil, nil
}

func GetByID(id int) (ILoanSessionEntity, error) {
	rows, err := storage.Db.Query("SELECT * FROM LoanSession WHERE ID = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	datas := []ILoanSessionEntity{}
	defer rows.Close()
	for rows.Next() {
		var ID, RequestLimit int
		var Date time.Time

		err := rows.Scan(&ID, &Date, &RequestLimit)
		if err != nil {
			return nil, err
		}

		var ls ILoanSessionEntity = loanSessionEntity{ID: ID, Date: Date, RequestLimit: RequestLimit}
		datas = append(datas, ls)
	}

	if len(datas) > 0 {
		return datas[0], nil
	}
	return nil, nil
}

func insertSession(limit int) error {
	if stmt, err := storage.Db.Prepare(`INSERT INTO LoanSession (Date, RequestLimit) VALUES (?, ?)`); err == nil {
		if _, err = stmt.Exec(getTodayDate(), limit); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func getTodayDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}
