package structs

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

// BindRequest will scann your struct and bind the request Values
// into your struct according to `json` tag on struct.
func BindRequest(values url.Values, items interface{}) error {
	val := reflect.ValueOf(items)

	if val.Kind() != reflect.Ptr {
		return errors.New("Target can't be value")
	}
	val = val.Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("json")
		t := typeField.Type.String()
		if len(values[tag]) > 0 {
			switch t {
			case "string":
				val.Field(i).SetString(values[tag][0])
			case "int64":
				res, _ := strconv.ParseInt(values[tag][0], 10, 64)
				val.Field(i).SetInt(res)
			case "bool":
				b := false
				if values[tag][0] == "1" {
					b = true
				}
				val.Field(i).SetBool(b)
			}
		}
	}
	return nil
}

// ValidateStruct will validate struct if `required` tag is equal to true.
func ValidateStruct(c interface{}) error {

	val := reflect.ValueOf(c).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("json")
		required := typeField.Tag.Get("required")
		t := typeField.Type.String()
		if required != "true" {
			continue
		}
		switch t {
		case "string":
			if val.Field(i).Interface().(string) == "" {
				return errors.New(tag + " is required.")
			}
		case "int64":
			if val.Field(i).Interface().(int64) == 0 {
				return errors.New(tag + " is required.")
			}
		}

	}
	return nil
}
