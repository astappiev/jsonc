# jsonc

[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/astappiev/jsonc) 

jsonc is a Go package that converts the jsonc format to standard json.

The jsonc format is like standard json but with following additional features:
- comments, single line (`// text`) or multiline (`/* text */`)
- trailing commas
- replace new line (`\n`) chars with `\\` following by `n`
- replace tabs (`\t`) chars with `\\` following by `t`
- remove invalid escape character of `'`
- escape `"` in string literals
- auto add missing commas

```text
{
  /* Dev Machine */
  "dbInfo": {
    "host": "localhost",
    "port": 5432,          
    "username": "josh",
    "password": "pass123", // please use a hashed password
  },

  "title": "Hello 
    World",
}
```

## Getting Started

### Installing

To start using jsonc, install Go and run `go get`:

```sh
$ go get -u github.com/astappiev/jsonc
```

This will retrieve the library.

### Example

There's a provided function `jsonc.ToJSON`, which does the conversion.

The resulting JSON will always be the same length as the input, and it will
include all the same line breaks at matching offsets. This is to ensure
the result can be later processed by an external parser and that that
parser will report messages or errors with the correct offsets.

The following example uses a JSON document that has comments and trailing
commas and converts it just prior to unmarshalling with the standard Go
JSON library.

```go
data := `
{
  /* Dev Machine */
  "dbInfo": {
    "host": "localhost",
    "port": 5432,          // use full email address
    "username": "josh",
    "password": "pass123", // use a hashed password
  },
  "title": "Hello 
    World",
}
`

err := json.Unmarshal(jsonc.ToJSON(data), &config)
```

### Performance

It's fast and can convert GB/s of jsonc to json.

## License

jsonc source code is available under the MIT [License](/LICENSE).
