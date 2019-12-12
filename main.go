package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dhpsagala/loan_process_donnysagala/customer"
	"github.com/dhpsagala/loan_process_donnysagala/loan"
	"github.com/dhpsagala/loan_process_donnysagala/loanSession"
	"github.com/dhpsagala/loan_process_donnysagala/province"
	rt "github.com/dhpsagala/loan_process_donnysagala/router"
	"github.com/dhpsagala/loan_process_donnysagala/storage"
)

var (
	session loanSession.ILoanSessionEntity
)

func main() {
	var err error
	appRtr := buildRoutes()

	storage.Init()
	storage.BuildScheme()

	fmt.Println(storage.Config)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to loan system")
	fmt.Println("-----------------------")

	if session, err = loanSession.GetCurrent(); err != nil {
		fmt.Println(err.Error())
		return
	}

	if session == nil {
		fmt.Println("No active session")
	}

	for {
		fmt.Print("-> ")

		if cmd, err := reader.ReadString('\n'); err == nil {
			cmdParts := strings.Fields(cmd)
			if cmdParts[0] == "exit" {
				break
			}
			appRtr.Listen(cmd)
		}
	}
}

func create_day_max(args ...interface{}) {
	var err error
	limit := args[0].(int)
	if session, err = loanSession.Load(limit); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("Created max request with %d requests", limit))
	}
}

func add(args ...interface{}) {
	var err error
	if session.CountLoan() < session.GetRequestLimit() {
		nik := args[0].(int)
		name := args[1].(string)
		age := args[2].(int)
		amount := args[3].(int)

		inValidCount := 0
		if !(amount >= storage.Config.AmountLimit.Min && amount <= storage.Config.AmountLimit.Max) {
			inValidCount++
		}

		if (amount % storage.Config.AmountLimit.Multiplier) != 0 {
			inValidCount++
		}

		if age < storage.Config.AgeLimit.Min || age > storage.Config.AgeLimit.Max {
			inValidCount++
		}

		var p province.IProvinceEntity
		if p, err = province.SearchByNIK(nik); err != nil {
			fmt.Println(err)
		} else {
			if !p.IsLoanAllowed() {
				inValidCount++
			}

			var c customer.ICustomerEntity
			if c, err = customer.GetByNIK(nik); err != nil {
				fmt.Println(err)
			} else {
				if c, err = customer.New(name, nik, age, p.GetID()); err != nil {
					fmt.Println(err)
				} else {
					isApproved := false
					if inValidCount == 0 {
						isApproved = true
					}
					var l loan.ILoanEntity
					var lID int
					if lID, err = session.GenerateNewLoanID(); err != nil {
						fmt.Println(err)
					} else {
						if l, err = loan.New(lID, session.GetID(), c.GetID(), amount, isApproved); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(fmt.Sprintf("Success: %d", l.GetLoanID()))
						}
					}
				}
			}
		}
	}
}

func status(args ...interface{}) {
	var err error
	var l loan.ILoanEntity
	if l, err = loan.GetByLoanID(args[0].(int)); err != nil {
		fmt.Println(err)
	} else {
		if l != nil {
			fmt.Println(l.StatusText())
		} else {
			fmt.Println("Sorry, doesn’t found it")
		}
	}
}

func installment(args ...interface{}) {
	var err error
	var l loan.ILoanEntity

	if l, err = loan.GetByLoanID(args[0].(int)); err != nil {
		fmt.Println(err)
	} else {
		if l == nil {
			fmt.Println("Sorry, doesn’t found it")
		} else {
			var s loanSession.ILoanSessionEntity
			if s, err = loanSession.GetByID(l.GetSessionID()); err != nil {
				fmt.Println(err)
			} else {
				monthLen := args[1].(int)
				adminFee := 100000 / monthLen
				capital := l.GetAmount() / monthLen
				total := adminFee + capital

				fmt.Printf("\n %s\t%s\t\t%s\t%s\t%s\t", "Month", "DueDate", "AdminFee", "Capital", "Total")
				month := 1
				for month <= monthLen {
					dueDate := s.GetDate().AddDate(0, month, 0)
					fmt.Printf("\n %d\t%s\t%d\t\t%d\t%d\t", month, dueDate.Format("02012006"), adminFee, capital, total)
					month++
				}
				fmt.Println("")
			}
		}
	}
}

func find_by_amount_accepted(args ...interface{}) {
	var lList []loan.ILoanEntity
	var err error

	if lList, err = loan.GetBySessionIdAndAmountAndStatus(session.GetID(), args[0].(int), true); err != nil {
		fmt.Println("Sorry, doesn’t found it")
	} else {
		lIDs := []int{}
		for _, l := range lList {
			lIDs = append(lIDs, l.GetID())
		}
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(lIDs), ",", " ", -1), "[]"))
	}
}

func find_by_amount_rejected(args ...interface{}) {
	var lList []loan.ILoanEntity
	var err error

	if lList, err = loan.GetBySessionIdAndAmountAndStatus(session.GetID(), args[0].(int), false); err != nil {
		fmt.Println("Sorry, doesn’t found it")
	} else {
		lIDs := []int{}
		for _, l := range lList {
			lIDs = append(lIDs, l.GetID())
		}
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(lIDs), ",", " ", -1), "[]"))
	}
}

func buildRoutes() rt.IRouter {
	r := rt.NewRouter()
	r.AddRoute("create_day_max", create_day_max, []string{"int"})
	r.AddRoute("add", add, []string{"int", "string", "int", "int"})
	r.AddRoute("status", status, []string{"int"})
	r.AddRoute("installment", installment, []string{"int", "int"})
	r.AddRoute("find_by_amount_accepted", find_by_amount_accepted, []string{"int"})
	r.AddRoute("find_by_amount_rejected", find_by_amount_rejected, []string{"int"})
	return r
}
