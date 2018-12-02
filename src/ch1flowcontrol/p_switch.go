package ch1flowcontrol

import (
	//"fmt"
	"reflect"
)

//func SwitchType(i interface{}) interface{}{
//	switch i.(type){
//	case bool: fmt.Println("Boolean")
//	case int: fmt.Println("Integer")
//	default: fmt.Println("String")
//	}
//	return i.(type)
//}

func ReflectType(i interface{}) interface{}{
	return reflect.TypeOf(i)
}
