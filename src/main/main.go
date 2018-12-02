package main

import (
	"fmt"
	"ch1flowcontrol"
)

const T = iota

func main(){
	fmt.Println("Hello Go")
	ch1flowcontrol.ReflectType2(12)
}