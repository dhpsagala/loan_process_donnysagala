package customer

type ICustomerEntity interface {
	GetID() int
}

type customerEntity struct {
	ID         int
	Name       string
	NIK        int
	Age        int
	ProvinceID int
}

func (c customerEntity) GetID() int {
	return c.ID
}
