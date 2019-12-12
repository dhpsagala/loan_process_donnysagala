package exception

import "fmt"

type NullArgumentLength struct {
}

func (e *NullArgumentLength) Error() string {
	return fmt.Sprintf("Expected one or more arguments.")
}
