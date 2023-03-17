# Change tag name to another
[Back to README.md](../README.md)

---
This will help when `default` tag name is used by another package. Or you just don't want to use tag `default`.

## How to archive that
Using `SetExtractTag`

```go
package main

import defaults "github.com/wizk3y/default-params"

type Params struct {
    SomeField string `custom:"some_value" json:"some_field"`
}

func init() {
    defaults.SetExtractTag("custom")
}

func main() {
    params := Params{}

    defaults.MustFillDefaultValue(&params)

    // params now is `{"some_field":"some_value"}`
}
```

---
[Back to README.md](../README.md)