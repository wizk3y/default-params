package internal

import (
	"fmt"
	"reflect"
)

var TypeAssertMap map[reflect.Kind]map[string]TypeAssert

func init() {
	TypeAssertMap = map[reflect.Kind]map[string]TypeAssert{
		reflect.String:  make(map[string]TypeAssert),
		reflect.Bool:    make(map[string]TypeAssert),
		reflect.Float64: make(map[string]TypeAssert),
		reflect.Float32: make(map[string]TypeAssert),
		reflect.Int:     make(map[string]TypeAssert),
		reflect.Int8:    make(map[string]TypeAssert),
		reflect.Int16:   make(map[string]TypeAssert),
		reflect.Int32:   make(map[string]TypeAssert),
		reflect.Int64:   make(map[string]TypeAssert),
	}
}

type TypeAssert interface {
	ToType(strValue string) interface{}
	ToPtrType(strValue string) interface{}
}

func getAssert(vType reflect.Type) TypeAssert {
	mapPkgAssert, exist := TypeAssertMap[vType.Kind()]
	if !exist {
		return nil
	}

	if vType.PkgPath() == "" {
		return nil
	}

	typeFullname := fmt.Sprintf("%s.%s", vType.PkgPath(), vType.Name())

	assert, exist := mapPkgAssert[typeFullname]
	if !exist {
		return nil
	}

	return assert
}
