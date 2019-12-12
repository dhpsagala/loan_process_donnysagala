package province

type IProvinceEntity interface {
	GetID() int
	IsLoanAllowed() bool
}

type provinceEntity struct {
	ID        int
	Name      string
	Code      int
	AllowLoan bool
}

func (p provinceEntity) GetID() int {
	return p.ID
}

func (p provinceEntity) IsLoanAllowed() bool {
	return p.AllowLoan
}
