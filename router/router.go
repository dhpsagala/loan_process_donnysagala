package router

import (
	"reflect"
	"strconv"
	"strings"

	ex "github.com/dhpsagala/loan_process_donnysagala/exception"
)

type IRouter interface {
	AddRoute(name string, handler func(args ...interface{}), paramsTypeMap []string)
	Listen(cmd string) error
}

type routerEntity struct {
	Routes map[string]*routeEntity
}

type routeEntity struct {
	handler       func(args ...interface{})
	paramsTypeMap []string
}

type routeHandler func(args ...interface{})

func NewRouter() IRouter {
	var r IRouter = &routerEntity{Routes: make(map[string]*routeEntity)}
	return r
}

func (r routerEntity) AddRoute(name string, handler func(args ...interface{}), paramsTypeMap []string) {
	if _, ok := r.Routes[name]; !ok {
		r.Routes[name] = &routeEntity{handler: handler, paramsTypeMap: paramsTypeMap}
	}
}

func (r routerEntity) Listen(cmd string) error {
	cmdParts := strings.Fields(cmd)

	if len(cmdParts) == 0 {
		return &ex.NullArgumentLength{}
	}

	routeName := cmdParts[0]
	params := removeItemOfStringArray(cmdParts, 0)

	if route, ok := r.Routes[routeName]; ok {
		if len(params) != len(route.paramsTypeMap) {
			return &ex.InvalidArgumentLength{ArgsLen: len(route.paramsTypeMap)}
		}
		fn := reflect.ValueOf(route.handler)
		args, err := buildArgs(params, route.paramsTypeMap)
		if err != nil {
			return err
		}
		fn.Call(args)
	} else {
		return &ex.RouteNotFound{Name: routeName}
	}
	return nil
}

func buildArgs(params []string, paramsTypeMap []string) ([]reflect.Value, error) {
	args := []reflect.Value{}
	for i, v := range params {
		paramType := paramsTypeMap[i]
		switch paramType {
		case "int":
			if intV, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				args = append(args, reflect.ValueOf(intV))
			}
		case "string":
			args = append(args, reflect.ValueOf(v))
		default:
			args = append(args, reflect.ValueOf(v))
		}
	}
	return args, nil
}

func removeItemOfStringArray(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
