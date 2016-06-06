package structs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const (
	ok        = "true"
	statePost = "POST"
	stateGet  = "GET"
)

// BindRequest will scan your struct and bind the request Values / Body
// into your struct according to `json` tag on struct.
func BindRequest(request *http.Request, target interface{}) error {
	if request.Method == statePost {
		body, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(body, &target)
		return nil
	}

	request.ParseForm()
	values := request.Form

	val := reflect.ValueOf(target)

	if val.Kind() != reflect.Ptr {
		return errors.New("Target can't be value")
	}
	val = val.Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("json")
		t := typeField.Type.String()
		if len(values[tag]) > 0 {
			values[tag][0] = strings.TrimSpace(values[tag][0])
			switch t {
			case "string":
				val.Field(i).SetString(values[tag][0])
			case "int64":
				r, _ := strconv.ParseFloat(values[tag][0], 64)
				val.Field(i).SetInt(int64(r))
			case "int32":
				r, _ := strconv.ParseFloat(values[tag][0], 64)
				val.Field(i).SetInt(int64(r))
			case "int16":
				r, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetInt(int64(r))
			case "int8":
				r, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetInt(int64(r))
			case "int":
				r, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetInt(int64(r))
			case "float32":
				res, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetFloat(res)
			case "float64":
				res, _ := strconv.ParseFloat(values[tag][0], 64)
				val.Field(i).SetFloat(res)
			case "bool":
				b := false
				if values[tag][0] == "1" || values[tag][0] == ok {
					b = true
				}
				val.Field(i).SetBool(b)
			default:
				return errors.New(t + " type is not supported. You can skip this binding by changing json tag value to `-`")
			}
		}
	}
	return nil
}

// ValidateStruct will validate struct if `required` tag is equal to true.
func ValidateStruct(target interface{}) error {
	val := reflect.ValueOf(target)

	if val.Kind() != reflect.Ptr {
		return errors.New("Target can't be value")
	}
	val = val.Elem()

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
		case "int":
			if val.Field(i).Interface().(int) == 0 {
				return errors.New(tag + " is required.")
			}
		case "int8":
			if val.Field(i).Interface().(int8) == 0 {
				return errors.New(tag + " is required.")
			}
		case "int16":
			if val.Field(i).Interface().(int16) == 0 {
				return errors.New(tag + " is required.")
			}
		case "int32":
			if val.Field(i).Interface().(int32) == 0 {
				return errors.New(tag + " is required.")
			}
		case "int64":
			if val.Field(i).Interface().(int64) == 0 {
				return errors.New(tag + " is required.")
			}
		case "float32":
			if val.Field(i).Interface().(float32) == 0 {
				return errors.New(tag + " is required.")
			}
		case "float64":
			if val.Field(i).Interface().(float64) == 0 {
				return errors.New(tag + " is required.")
			}
		}

	}
	return nil
}
