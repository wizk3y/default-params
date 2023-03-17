package defaults

import (
	"reflect"

	"github.com/wizk3y/default-params/assert"
	"github.com/wizk3y/default-params/internal"
)

func init() {
	RegisterAssertFunc(reflect.Int64, "time.Duration", assert.NewDurationTypeAssert("ns"))
}

func RegisterAssertFunc(kind reflect.Kind, pkgName string, typeAssert internal.TypeAssert) {
	if _, ok := internal.TypeAssertMap[kind]; !ok {
		internal.TypeAssertMap[kind] = make(map[string]internal.TypeAssert)
	}

	internal.TypeAssertMap[kind][pkgName] = typeAssert
}
