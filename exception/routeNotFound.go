package exception

import "fmt"

type RouteNotFound struct {
	Name string
}

func (e *RouteNotFound) Error() string {
	return fmt.Sprintf("Route %s is not found", e.Name)
}
