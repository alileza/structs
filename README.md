# Structs [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/alileza/structs) ![CircleCI](https://circleci.com/gh/alileza/structs.png?style=shield&circle-token=e9a794ff32b429804e757d05de595d32fdbd1929) [![Coverage Status](https://coveralls.io/repos/github/alileza/structs/badge.svg?branch=master)](https://coveralls.io/github/alileza/structs?branch=master)

This library helps you to play around with structs, such as [bind incoming request into struct](#bind-request), validate struct, etc.

### Bind Request
BindRequest will scan your struct and bind the request Values / Body into your struct according to `json` tag on struct.

[Example Here](https://godoc.org/github.com/alileza/structs#example-BindRequest)

### Validate Struct
ValidateStruct will validate struct if `required` tag is equal to true.

[Example Here](https://godoc.org/github.com/alileza/structs#example-ValidateStruct)

### To Map
ToMap returns map following the input struct.

[Example Here](https://godoc.org/github.com/alileza/structs#example-BindRequest)
