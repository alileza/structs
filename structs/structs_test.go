package structs

import (
	"net/url"
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
	TSkip        string     `json:"t_skip" skip:"true"`
}

type testCase struct {
	Key    string
	Input  string
	Output interface{}
}

func TestBindRequest(t *testing.T) {

	var target bindRequestStruct
	values := url.Values{}
	values.Add("t_int", "123")
	BindRequest(values, &target)
	if target.TInt != 123 {
		t.Error("t_int mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int8", "123")
	BindRequest(values, &target)
	if target.TInt8 != 123 {
		t.Error("t_int8 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int16", "123")
	BindRequest(values, &target)
	if target.TInt16 != 123 {
		t.Error("t_int16 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int32", "123")
	BindRequest(values, &target)
	if target.TInt32 != 123 {
		t.Error("t_int32 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int64", "123")
	BindRequest(values, &target)
	if target.TInt64 != 123 {
		t.Error("t_int64 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_float32", "123.221")
	BindRequest(values, &target)
	if target.TFloat32 != 123.221 {
		t.Error("t_float32 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_float64", "123.223841")
	BindRequest(values, &target)
	if target.TFloat64 != 123.223841 {
		t.Error("t_float64 mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_string", "123.223841")
	BindRequest(values, &target)
	if target.TString != "123.223841" {
		t.Error("t_string mismatch !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "true")
	BindRequest(values, &target)
	if target.TBool != true {
		t.Error("t_bool mismatch (true)!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "1")
	BindRequest(values, &target)
	if target.TBool != true {
		t.Error("t_bool mismatch (1)!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "false")
	BindRequest(values, &target)
	if target.TBool != false {
		t.Error("t_bool mismatch (false)!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_bool", "")
	BindRequest(values, &target)
	if target.TBool != false {
		t.Error("t_bool mismatch ('')!")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_unsupported", "123.223841")
	err := BindRequest(values, &target)
	if err.Error() != "complex128 type is not supported. You can skip this binding by changing json tag value to `-`" {
		t.Error("Unsupported not working !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_skip", "123.223841")
	err = BindRequest(values, &target)
	if target.TSkip == "123.223841" {
		t.Error("skip not working !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int8", "123.223841")
	BindRequest(values, &target)
	if target.TInt8 != 123 {
		t.Error("float to int fail !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_int64", "123192733.22338723841")
	BindRequest(values, &target)
	if target.TInt64 != 123192733 {
		t.Error("float to int fail !")
	}

	target = bindRequestStruct{}
	values = url.Values{}
	values.Add("t_float64", "123192733")
	BindRequest(values, &target)
	if target.TFloat64 != 123192733 {
		t.Error("int to float fail !")
	}

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
