# default-params

default-params is a Go library help you fill default value into struct, which helpful when writing func with many args just using Go stdlib.

## Install
```shell
go get github.com/wizk3y/default-params
```

**Note:** default-params uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies.

## Usage
- Create struct with tag `default`
```go
// struct.go
package main

type Params struct {
    SomeField string `default:"default_value"`
}
```
- Using `FillDefaultValue` or `MustFillDefaultValue`
```go
// main.go
package main

import defaults "github.com/wizk3y/default-params"

func main() {
    params := Params{}

    defaults.MustFillDefaultValue(&params)

    // params now is `{"some_field":"some_value"}`
}
```

## Advance usage
- [Change tag name to another](docs/name_tag.md)
- [Nil value for pointer struct](docs/nil_struct.md)
- [Extend with custom type base on primitive type](docs/extend.md)
- [Override setted value or struct](docs/override.md)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)