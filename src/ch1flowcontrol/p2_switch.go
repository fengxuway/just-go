package ch1flowcontrol

import (
	//"fmt"
	"reflect"
	"ch1flowcontrol/sub"
	"fmt"
)

//func SwitchType(i interface{}) interface{}{
//	switch i.(type){
//	case bool: fmt.Println("Boolean")
//	case int: fmt.Println("Integer")
//	default: fmt.Println("String")
//	}
//	return i.(type)
//}

func ReflectType2(i interface{}) interface{}{
	fmt.Println(sub.Sub())
	return reflect.TypeOf(i)
}
