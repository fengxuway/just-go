package ch1flowcontrol

import(
	"testing"
	"fmt"
)

func Test_Reflect_Type(t *testing.T){
	fmt.Println("Test Type: ", ReflectType(ReflectType(12)))
	i, j := 100,12
	if i < j{
		t.Error("ERROR: i < j ")
	}else{
		t.Log("Success: i > j")
	}
}
