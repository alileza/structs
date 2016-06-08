# Structs [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/alileza/structs) ![CircleCI](https://circleci.com/gh/alileza/structs.png?style=shield&circle-token=e9a794ff32b429804e757d05de595d32fdbd1929) [![Coverage Status](https://coveralls.io/repos/github/alileza/structs/badge.svg?branch=master)](https://coveralls.io/github/alileza/structs?branch=master)

This library helps you to play around with structs, such as [bind incoming request into struct](#bind-request), validate struct, etc.

### Bind Request
BindRequest will scan your struct and bind the request Values / Body into your struct according to `json` tag on struct.

##### Example
```
// Request example: http://localhost:9000?api_key=123456
values := url.Values{}
values.Add("api_key", "123456")
req := &http.Request{
    Form: values,
}

var target struct {
    APIKey string `json:"api_key"`
}
err := BindRequest(req, &target)
if err != nil {
    log.Fatal("request parsing, failed.")
}
fmt.Println(target.APIKey)
```
Output :
```
123456
```

### Validate Struct
ValidateStruct will validate struct if `required` tag is equal to true.

##### Example
```
MyStruct := struct {
    Name string `json:"name" required:"true"`
}{}

err := ValidateStruct(&MyStruct)

fmt.Println(err)
```
Output :
```
name is required.
```

### To Map
ToMap returns map following the input struct.

##### Example
```
MyStruct := struct {
    Name    string `json:"name"`
    Age     int64
    Address struct {
        Hometown string
    }
}{
    Name: "Arya Stark",
    Age:  14,
    Address: struct {
        Hometown string
    }{
        Hometown: "Winterfell",
    },
}

myMap := ToMap(MyStruct)

fmt.Println(myMap["name"])
fmt.Println(myMap["Age"])
fmt.Println(myMap["Address"].(map[string]interface{})["Hometown"])
```
Output:
```
Arya Stark
14
Winterfell
```

## Docs
[]
