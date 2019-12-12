package exception

import "fmt"

type InvalidArgumentLength struct {
	ArgsLen int
}

func (e *InvalidArgumentLength) Error() string {
	return fmt.Sprintf("Expected %d arguments.", e.ArgsLen)
}
