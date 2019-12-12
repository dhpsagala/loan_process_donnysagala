package customer

import (
	"database/sql"

	"github.com/dhpsagala/loan_process_donnysagala/storage"
)

func New(name string, nik int, age int, provinceID int) (ICustomerEntity, error) {
	var err error
	var stmt *sql.Stmt
	if stmt, err = storage.Db.Prepare(`INSERT OR IGNORE INTO Customer (Name, NIK, Age, ProvinceID) VALUES (?, ?, ?, ?)`); err == nil {
		if _, err = stmt.Exec(name, nik, age, provinceID); err == nil {
			return GetByNIK(nik)
		}
	}
	return nil, err
}

func GetByNIK(nik int) (ICustomerEntity, error) {
	var err error
	var rows *sql.Rows
	if rows, err = storage.Db.Query(`SELECT * FROM Customer WHERE NIK = ?`, nik); err == nil {
		datas := []ICustomerEntity{}
		defer rows.Close()
		for rows.Next() {
			var ID, NIK, Age, ProvinceID int
			var Name string

			if err = rows.Scan(&ID, &Name, &NIK, &Age, &ProvinceID); err == nil {
				var c ICustomerEntity = customerEntity{ID: ID, NIK: NIK, Age: Age, Name: Name, ProvinceID: ProvinceID}
				datas = append(datas, c)
			}
		}

		if len(datas) > 0 {
			return datas[0], nil
		}
	}
	return nil, err
}
