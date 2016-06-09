package structs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"testing"
)

type bindRequestStruct struct {
	TInt         int        `json:"t_int" required:"true"`
	TInt8        int8       `json:"t_int8" required:"true"`
	TInt16       int16      `json:"t_int16" required:"true"`
	TInt32       int32      `json:"t_int32" required:"true"`
	TInt64       int64      `json:"t_int64" required:"true"`
	TFloat32     float32    `json:"t_float32" required:"true"`
	TFloat64     float64    `json:"t_float64" required:"true"`
	TBool        bool       `json:"t_bool" required:"true"`
	TString      string     `json:"t_string" required:"true"`
	TUnsupported complex128 `json:"t_unsupported" required:"true"`
}

func ExampleBindRequest() {
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
	// Output: 123456
}

func TestBindRequest(t *testing.T) {
	req := &http.Request{}

	var target bindRequestStruct
	values := url.Values{}
	values.Add("t_int", "123")
	req.Form = values
	BindRequest(req, &target)
	if target.TInt != 123 {
		t.Error("t_int mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int8", "123")
	req.Form = values
	BindRequest(req, &target)
	if target.TInt8 != 123 {
		t.Error("t_int8 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int16", "123")
	req.Form = values
	BindRequest(req, &target)
	if target.TInt16 != 123 {
		t.Error("t_int16 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int32", "123")
	req.Form = values
	BindRequest(req, &target)
	if target.TInt32 != 123 {
		t.Error("t_int32 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int64", "123")
	req.Form = values
	BindRequest(req, &target)
	if target.TInt64 != 123 {
		t.Error("t_int64 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_float32", "123.221")
	req.Form = values
	BindRequest(req, &target)
	if target.TFloat32 != 123.221 {
		t.Error("t_float32 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_float64", "123.223841")
	req.Form = values
	BindRequest(req, &target)
	if target.TFloat64 != 123.223841 {
		t.Error("t_float64 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_string", "123.223841")
	req.Form = values
	BindRequest(req, &target)
	if target.TString != "123.223841" {
		t.Error("t_string mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "true")
	req.Form = values
	BindRequest(req, &target)
	if target.TBool != true {
		t.Error("t_bool mismatch (true)!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "1")
	req.Form = values
	BindRequest(req, &target)
	if target.TBool != true {
		t.Error("t_bool mismatch (1)!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "false")
	req.Form = values
	BindRequest(req, &target)
	if target.TBool != false {
		t.Error("t_bool mismatch (false)!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "")
	req.Form = values
	BindRequest(req, &target)
	if target.TBool != false {
		t.Error("t_bool mismatch ('')!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_unsupported", "123.223841")
	req.Form = values
	err := BindRequest(req, &target)
	if err.Error() != "complex128 type is not supported. You can skip this binding by changing json tag value to `-`" {
		t.Error("Unsupported not working !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int8", "123.223841")
	req.Form = values
	BindRequest(req, &target)
	if target.TInt8 != 123 {
		t.Error("float to int fail !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int64", "123192733.22338723841")
	req.Form = values
	BindRequest(req, &target)
	if target.TInt64 != 123192733 {
		t.Error("float to int fail !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_float64", "123192733")
	req.Form = values
	BindRequest(req, &target)
	if target.TFloat64 != 123192733 {
		t.Error("int to float fail !")
	}

	target = bindRequestStruct{}
	req.Method = "POST"
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(`{"t_int8" : 123 }`)))
	BindRequest(req, &target)
	if target.TInt8 != 123 {
		t.Error("post t_int8 read fail !")
	}

	target = bindRequestStruct{}
	req.Method = "POST"
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(`{"t_int64" : 93934123 }`)))
	BindRequest(req, &target)
	if target.TInt64 != 93934123 {
		t.Error("post t_int64 read fail !")
	}

	target = bindRequestStruct{}
	req.Method = "POST"
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(`{"t_bool" : true }`)))
	BindRequest(req, &target)
	if target.TBool != true {
		t.Error("post t_bool read fail !")
	}

	target = bindRequestStruct{}
	req.Method = "POST"
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(`{"t_bool" : false }`)))
	BindRequest(req, &target)
	if target.TBool != false {
		t.Error("post t_bool read fail !")
	}

	target = bindRequestStruct{}
	req.Method = "POST"
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(`{"t_float32" : 123.3223 }`)))
	BindRequest(req, &target)
	if target.TFloat32 != 123.3223 {
		t.Error("post t_float32 read fail !")
	}

	target = bindRequestStruct{}
	req.Method = "POST"
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(`{"t_string" : "123.3223" }`)))
	BindRequest(req, &target)
	if target.TString != "123.3223" {
		t.Error("post t_string read fail !")
	}

}

func ExampleValidateStruct() {

	MyStruct := struct {
		Name string `json:"name" required:"true"`
	}{}

	err := ValidateStruct(&MyStruct)

	fmt.Println(err)
	// Output: name is required.
}

func TestValidateStruct(t *testing.T) {

	var target bindRequestStruct
	err := ValidateStruct(&target)
	if err.Error() != "t_int is required." {
		t.Error("t_int validation fail !")
	}

	target.TInt = 1
	err = ValidateStruct(&target)
	if err.Error() != "t_int8 is required." {
		t.Error("t_int8 validation fail !")
	}

	target.TInt8 = 1
	err = ValidateStruct(&target)
	if err.Error() != "t_int16 is required." {
		t.Error("t_int16 validation fail !")
	}

	target.TInt16 = 1
	err = ValidateStruct(&target)
	if err.Error() != "t_int32 is required." {
		t.Error("t_int32 validation fail !")
	}

	target.TInt32 = 1
	err = ValidateStruct(&target)
	if err.Error() != "t_int64 is required." {
		t.Error("t_int64 validation fail !")
	}

	target.TInt64 = 1
	err = ValidateStruct(&target)
	if err.Error() != "t_float32 is required." {
		t.Error("t_float32 validation fail !")
	}

	target.TFloat32 = 1
	err = ValidateStruct(&target)
	if err.Error() != "t_float64 is required." {
		t.Error("t_float64 validation fail !")
	}

	target.TFloat64 = 1
	err = ValidateStruct(&target)
	if err.Error() != "t_string is required." {
		t.Error("t_string validation fail !")
	}
}

func ExampleToMap() {

	MyStruct := struct {
		Name    string `json:"name"`
		Age     int64
		Address struct {
			Hometown  string
			Latitude  float64
			Longitude float64
		}
	}{
		Name: "Arya Stark",
		Age:  14,
		Address: struct {
			Hometown  string
			Latitude  float64
			Longitude float64
		}{
			Hometown: "Winterfell",
			Latitude: 12.32,
		},
	}

	myMap := ToMap(MyStruct)
	myMapStr := ToMap(MyStruct, true)
	obj, _ := json.Marshal(myMapStr)

	fmt.Println(myMap["name"])
	fmt.Println(myMap["Age"])
	fmt.Println(myMap["Address"].(map[string]interface{})["Hometown"])
	fmt.Println(string(obj))
	// Output: Arya Stark
	// 14
	// Winterfell
	// {"Address":{"Hometown":"Winterfell","Latitude":"12.32","Longitude":"0.00"},"Age":"14","name":"Arya Stark"}

}

func TestToMap(t *testing.T) {
	testStruct := struct {
		Name       string `json:"name"`
		Age        int64
		FavNumbers []int
		Address    struct {
			Street string
		}
	}{
		Name:       "Ali",
		Age:        22,
		FavNumbers: []int{15, 3, 22},
	}
	testStruct.Address.Street = "flamboyan"

	result := ToMap(testStruct)

	if _, ok := result["name"]; !ok {
		t.Error("failed to use json tag!")
	}

	if val, ok := result["Age"]; !ok {
		t.Error("failed to convert to map!")
	} else if val.(int64) != 22 {
		t.Error("failed to convert to map!")
	}

	if val, ok := result["Address"]; !ok {
		t.Error("failed to convert multi-struct to map!")
	} else if val.(map[string]interface{})["Street"] != "flamboyan" {
		t.Error("failed to access multi-struct to map!")
	}

	if val, ok := result["FavNumbers"]; !ok {
		t.Error("failed to handle slice!")
	} else if reflect.ValueOf(val).Index(0).Interface() != 15 {
		t.Error("failed to handle slice!")
	} else if reflect.ValueOf(val).Index(2).Interface() != 22 {
		t.Error("failed to handle slice!")
	}
}

func TestToMapString(t *testing.T) {
	type userAddress struct {
		Street      string
		Number      int32
		FootSize    int
		Geolocation struct {
			Lat float64
			Lng float32
		}
	}
	testStruct := struct {
		Name       string `json:"name"`
		Age        int64
		FavNumbers interface{}
		Friends    []struct {
			Name int32
		}
		Address userAddress
	}{
		Friends: []struct {
			Name int32
		}{{
			Name: 54,
		}, {
			Name: 34,
		}},
		FavNumbers: []int{15, 3, 22},
		Name:       "Ali",
		Age:        22,
	}
	testStruct.Address.Street = "flamboyan"
	testStruct.Address.Number = 5
	testStruct.Address.FootSize = 12335
	testStruct.Address.Geolocation.Lat = 83.23
	testStruct.Address.Geolocation.Lng = 89.22

	result := ToMap(testStruct, true)
	b, _ := json.Marshal(result)
	fmt.Println(string(b))

	if _, ok := result["name"]; !ok {
		t.Error("failed to use json tag!")
	}

	if val, ok := result["Age"]; !ok {
		t.Error("failed to convert to map!")
	} else if val.(string) != "22" {
		t.Error("failed to convert to string!")
	}

	val, ok := result["Address"]

	if !ok {
		t.Error("failed to convert multi-struct to map!")
	}

	if val.(map[string]interface{})["Number"] != "5" {
		t.Error("failed to access multi-struct int to string!")
	}

	if val.(map[string]interface{})["Geolocation"].(map[string]interface{})["Lat"] != "83.23" {
		t.Error("failed to access double multi-struct float to string!")
	}

	if val.(map[string]interface{})["Geolocation"].(map[string]interface{})["Lng"] != "89.22" {
		t.Error("failed to access double multi-struct float to string!")
	}

	if val, ok := result["FavNumbers"]; !ok {
		t.Error("failed to handle slice!")
	} else if reflect.ValueOf(val).Index(0).Interface() != "15" {
		t.Error("failed to handle slice!")
	} else if reflect.ValueOf(val).Index(2).Interface() != "22" {
		t.Error("failed to handle slice!")
	}
}
