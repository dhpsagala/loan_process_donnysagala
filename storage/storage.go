package storage

import (
	"database/sql"
	"log"

	cfg "github.com/dhpsagala/loan_process_donnysagala/config"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Db     *sql.DB
	Config cfg.ConfigObj
)

func Init() {
	Config = cfg.Load()
	database, err := sql.Open("sqlite3", Config.Database.Path)
	handleError(err)
	Db = database
}

func BuildScheme() {
	createCustomerTable()
	createProvinceTable()
	createLoanSessionTable()
	createLoanTable()
	insertProvinceDatas()
}

func createCustomerTable() {
	statement, err := Db.Prepare(
		`CREATE TABLE IF NOT EXISTS Customer (
			ID INTEGER PRIMARY KEY,
			Name TEXT,
			NIK INTEGER UNIQUE,
			Age INTEGER,
			ProvinceID INTEGER NULL
		)`,
	)
	handleError(err)

	_, err = statement.Exec()
	handleError(err)
}

func createProvinceTable() {
	statement, err := Db.Prepare(
		`CREATE TABLE IF NOT EXISTS Province (
			ID INTEGER PRIMARY KEY,
			Name TEXT,
			Code INTEGER UNIQUE,
			AllowLoan BOOLEAN
		)`,
	)
	handleError(err)

	_, err = statement.Exec()
	handleError(err)
}

func createLoanSessionTable() {
	statement, err := Db.Prepare(
		`CREATE TABLE IF NOT EXISTS LoanSession (
			ID INTEGER PRIMARY KEY,
			Date DATETIME,
			RequestLimit INTEGER
		)`,
	)
	handleError(err)

	_, err = statement.Exec()
	handleError(err)
}

func createLoanTable() {
	statement, err := Db.Prepare(
		`CREATE TABLE IF NOT EXISTS Loan (
			ID INTEGER PRIMARY KEY,
			LoanID INTEGER,
			SessionID INTEGER,
			CustomerID INTEGER,
			Amount INTEGER,
			CreateDate DATETIME,
			IsApproved BOOLEAN
		)`,
	)
	handleError(err)

	_, err = statement.Exec()
	handleError(err)
}

func insertProvinceDatas() {
	query := `
		INSERT OR IGNORE INTO PROVINCE (Name, Code, AllowLoan) VALUES ('ACEH', '11', 0),
		('SUMATERA UTARA', '12', 1),
		('SUMATERA BARAT', '13', 0),
		('RIAU', '14', 0),
		('JAMBI', '15', 0),
		('SUMATERA SELATAN', '16', 0),
		('BENGKULU', '17', 0),
		('LAMPUNG', '18', 0),
		('KEPULAUAN BANGKA BELITUNG', '19', 0),
		('KEPULAUAN RIAU', '21', 0),
		('DKI JAKARTA', '31', 1),
		('JAWA BARAT', '32', 1),
		('JAWA TENGAH', '33', 0),
		('DI YOGYAKARTA', '34', 0),
		('JAWA TIMUR', '35', 1),
		('BANTEN', '36', 0),
		('BALI', '51', 0),
		('NUSA TENGGARA BARAT', '52', 0),
		('NUSA TENGGARA TIMUR', '53', 0),
		('KALIMANTAN BARAT', '61', 0),
		('KALIMANTAN TENGAH', '62', 0),
		('KALIMANTAN SELATAN', '63', 0),
		('KALIMANTAN TIMUR', '64', 0),
		('SULAWESI UTARA', '71', 0),
		('SULAWESI SELATAN', '73', 0),
		('SULAWESI TENGGARA', '74', 0),
		('GORONTALO', '75', 0),
		('SULAWESI BARAT', '76', 0),
		('MALUKU', '81', 0),
		('MALUKU UTARA', '82', 0),
		('PAPUA', '94', 0),
		('PAPUA BARAT', '91', 0),
		('SULAWESI TENGAH', '72', 0),
		('KALIMANTAN UTARA', '65', 0);
	`
	tx, err := Db.Begin()
	handleError(err)

	statement, err := Db.Prepare(query)
	handleError(err)

	_, err = statement.Exec()
	if err != nil {
		tx.Rollback()
		handleError(err)
	}

	handleError(tx.Commit())
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
