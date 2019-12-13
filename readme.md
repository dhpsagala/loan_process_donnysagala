# Requirement
- [Go](https://golang.org/) as programming language
- [sqlite3](https://www.sqlite.org/) as database

# How to run
Via Git
```bash
$ git clone https://github.com/dhpsagala/loan_process_donnysagala.git
$ cd loan_process_donnysagala
$ go run .\main.go
```

Via .zip file
- download .zip from [here](https://github.com/dhpsagala/loan_process_donnysagala/archive/master.zip)
- extract .zip to a folder
- open terminal/cmd in the folder
- run with this command
```bash
$ go run .\main.go
$ loan_process_donnysagala.exe
```

if you are using windows then you can follow this step
- open terminal/cmd in the app folder
- type this command
```bash
$ go build
$ loan_process_donnysagala.exe
```

# Usage
- Create current day session limit
```bash
-> create_day_max {limit}
```
- Create loan for customer
```bash
-> add {id_card_number} {customer_name} {customer_age} {loan_amount}
```
- Check loan status
```bash
-> status {loan_id}
```
- Generate loan installment by given month range
```bash
-> installment {loan_id} {month_range}
```
- find every accepted loan id by given amount
```bash
-> find_by_amount_accepted {loan_amount}
```
- find every rejected loan id by given amount
```bash
-> find_by_amount_rejected {loan_amount}
```
- exit app
```bash
-> exit
```