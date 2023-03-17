# Nil value for pointer struct
[Back to README.md](../README.md)

---
Even default value of pointer is struct, but this library iterate every field and go deep down to fill value.

```go
package main

import defaults "github.com/wizk3y/default-params"

type SubParams struct {
    SubField string `default:"sub_value" json:"sub_field"`
}

type Params struct {
    Sub *SubField `json:"sub"`
}

func main() {
    params := Params{}

    defaults.MustFillDefaultValue(&params)

    // params now is `{"sub":{"sub_field":"sub_value"}}`
}
```

## How to prevent that happen
Simple add nil sign with `default` tag to pointer struct. Default nil sign is `-`
```go
package main

import defaults "github.com/wizk3y/default-params"

type SubParams struct {
    SubField string `default:"sub_value" json:"sub_field"`
}

type Params struct {
    Sub *SubField `default:"-" json:"sub"`
}

func main() {
    params := Params{}

    defaults.MustFillDefaultValue(&params)

    // params now is `{"sub":null}`
}
```

**Note:** You can also change nil sign to another character set with `SetNilSign`
```go
package main

import defaults "github.com/wizk3y/default-params"

type SubParams struct {
    SubField string `default:"sub_value" json:"sub_field"`
}

type Params struct {
    Sub *SubField `default:"nil" json:"sub"`
}

func init() {
    defaults.SetNilSign("nil")
}

func main() {
    params := Params{}

    defaults.MustFillDefaultValue(&params)

    // params now is `{"sub":null}`
}
```

---
[Back to README.md](../README.md)