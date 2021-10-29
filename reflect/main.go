package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data interface{}
	data = map[string]string{
		"Nama": "Ipan",
	}
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Map {
		for _, key := range rv.MapKeys() {
			strct := rv.MapIndex(key)
			fmt.Println(key.Interface(), strct.Interface())
		}
	}
	if rv.Kind() == reflect.Struct {
	}
}
