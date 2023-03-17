# Extend with custom type base on primitive type
    [Back to README.md](../README.md)

---
You can also control how value to fill to custom type (base on primitive type) by implement a struct with 2 func `ToType` and `ToPtrType`. See best practice in [extend_test.go](../extend_test.go)
```go
type Sample string

func samplePtr(s Sample) *Sample {
	return &s
}

type sampleTypeAssert struct{}

func (a *sampleTypeAssert) ToType(strValue string) interface{} {
	return Sample(fmt.Sprintf("sample-%s", strValue))
}

func (a *sampleTypeAssert) ToPtrType(strValue string) interface{} {
	return samplePtr(a.ToType(strValue).(Sample))
}
```

**Note 1:** this library default come with `time.Duration` - which base on `int64` primitive type - assert and auto add suffix `ns`. Eg: default value `1000000000` will be parse by `time.ParseDuration` with args is `1000000000ns`. 

**Note 2:** You can override default suffix above with running `RegisterAssertFunc` as below. Accepted values is `h`, `m`, `s`, `ms`, `µs`, `us`, `ns` due to `time.ParseDuration` valid time units.
```go
import (
    defaults "github.com/wizk3y/default-params"
    "github.com/wizk3y/default-params/assert"
)

func init() {
    defaults.RegisterAssertFunc(
        reflect.Int64, 
        "time.Duration", 
        assert.NewDurationTypeAssert("s"),
    )
}
```

---
[Back to README.md](../README.md)