package utils

import (
	"reflect"
)

// El primer valor es el valor que queremos setear, el segundo es el valor desde donde queremos setear
func SetIfNotNil(target interface{}, value interface{}) {
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		reflect.ValueOf(target).Elem().Set(val.Elem())
	}
}

// Called on string == ""
func VerifyNotNullOrEmpty(s string, lambda func()) {
	if s == "" {
		lambda()
	}
}
