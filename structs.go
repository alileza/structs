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
	ok           = "true"
	statePost    = "POST"
	stateGet     = "GET"
	stateInt     = "int"
	stateInt8    = "int8"
	stateInt16   = "int16"
	stateInt32   = "int32"
	stateInt64   = "int64"
	stateFloat32 = "float32"
	stateFloat64 = "float64"
	stateBool    = "bool"
	stateString  = "string"
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
			case stateString:
				val.Field(i).SetString(values[tag][0])
			case stateInt64:
				r, _ := strconv.ParseFloat(values[tag][0], 64)
				val.Field(i).SetInt(int64(r))
			case stateInt32:
				r, _ := strconv.ParseFloat(values[tag][0], 64)
				val.Field(i).SetInt(int64(r))
			case stateInt16:
				r, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetInt(int64(r))
			case stateInt8:
				r, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetInt(int64(r))
			case stateInt:
				r, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetInt(int64(r))
			case stateFloat32:
				res, _ := strconv.ParseFloat(values[tag][0], 32)
				val.Field(i).SetFloat(res)
			case stateFloat64:
				res, _ := strconv.ParseFloat(values[tag][0], 64)
				val.Field(i).SetFloat(res)
			case stateBool:
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
		case stateString:
			if val.Field(i).Interface().(string) == "" {
				return errors.New(tag + " is required.")
			}
		case stateInt:
			if val.Field(i).Interface().(int) == 0 {
				return errors.New(tag + " is required.")
			}
		case stateInt8:
			if val.Field(i).Interface().(int8) == 0 {
				return errors.New(tag + " is required.")
			}
		case stateInt16:
			if val.Field(i).Interface().(int16) == 0 {
				return errors.New(tag + " is required.")
			}
		case stateInt32:
			if val.Field(i).Interface().(int32) == 0 {
				return errors.New(tag + " is required.")
			}
		case stateInt64:
			if val.Field(i).Interface().(int64) == 0 {
				return errors.New(tag + " is required.")
			}
		case stateFloat32:
			if val.Field(i).Interface().(float32) == 0 {
				return errors.New(tag + " is required.")
			}
		case stateFloat64:
			if val.Field(i).Interface().(float64) == 0 {
				return errors.New(tag + " is required.")
			}
		}

	}
	return nil
}

// ToMap returns map following the input struct.
// Second params is used to conver map values into string.
func ToMap(target interface{}, opts ...bool) map[string]interface{} {
	var (
		key string
		dt  bool
	)
	if len(opts) > 0 {
		dt = opts[0]
	}
	result := make(map[string]interface{}, 1)
	v := reflect.ValueOf(target)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("json")
		vval := v.Field(i)

		if tag == "-" || !vval.CanInterface() {
			continue
		} else if tag != "" {
			key = tag
		} else {
			key = v.Type().Field(i).Name
		}

		value := vval.Interface()
		if value == nil {
			value = nil
		}

		if reflect.TypeOf(value).Kind() == reflect.Slice {
			val := reflect.ValueOf(value)
			tmp := make([]interface{}, val.Len())

			for i := 0; i < val.Len(); i++ {
				t := val.Index(i).Interface()
				typ := reflect.TypeOf(t)
				if typ.Name() == "" {
					tmp[i] = ToMap(t, dt)
				} else if dt {
					tmp[i] = toString(t)
				} else {
					tmp[i] = t
				}
			}
			result[key] = tmp
		} else if reflect.TypeOf(value).Kind() == reflect.Struct {
			result[key] = ToMap(value, dt)
		} else if dt {
			result[key] = toString(value)
		} else {
			result[key] = value
		}
	}

	return result
}

func toString(v interface{}) interface{} {

	if reflect.TypeOf(v).Name() == stateInt {
		return strconv.Itoa(v.(int))
	} else if reflect.TypeOf(v).Name() == stateInt8 {
		return strconv.Itoa(int(v.(int8)))
	} else if reflect.TypeOf(v).Name() == stateInt16 {
		return strconv.Itoa(int(v.(int16)))
	} else if reflect.TypeOf(v).Name() == stateInt32 {
		return strconv.Itoa(int(v.(int32)))
	} else if reflect.TypeOf(v).Name() == stateInt64 {
		return strconv.Itoa(int(v.(int64)))
	} else if reflect.TypeOf(v).Name() == stateFloat32 {
		return strconv.FormatFloat(float64(v.(float32)), 'f', 2, 32)
	} else if reflect.TypeOf(v).Name() == stateFloat64 {
		return strconv.FormatFloat(v.(float64), 'f', 2, 64)
	} else if reflect.TypeOf(v).Name() == stateBool {
		return v.(bool)
	}
	return v
}

func Copy(from interface{}, target interface{}) error {

	if reflect.ValueOf(target).Kind() != reflect.Ptr {
		return errors.New("Target must be a pointer.")
	}

	tmp, err := json.Marshal(from)
	if err != nil {
		return err
	}
	json.Unmarshal(tmp, &target)
	return nil
}
