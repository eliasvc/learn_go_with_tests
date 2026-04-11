package main

import "reflect"

func walk(x any, fn func(string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walk(val.MapIndex(k).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}
	}
}
