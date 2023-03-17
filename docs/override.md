# Override setted value or struct
[Back to README.md](../README.md)

---
This library also support for use case *reset* struct to some default state.

## How to archive that
By passing `OverrideSettedValueOpt` when calling `FillDefaultValue` or `MustFillDefaultValue`
```go
package main

import defaults "github.com/wizk3y/default-params"

type Params struct {
    SomeField string `default:"default_value"`
}

func main() {
    params := Params{
        SomeField: "setted_value",
    }

    defaults.MustFillDefaultValue(&params, defaults.OverrideSettedValueOpt())

    // params now is `{"some_field":"some_value"}`
}
```

---
[Back to README.md](../README.md)