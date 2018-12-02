package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Height int
	Weight int
}

func (p Person) Method (name string) string {
	s := "Name is " + name + " Height: " + strconv.Itoa(p.Height)
	fmt.Println(s)
	return s
}

func main(){
	p := Person{175, 60}
	v := reflect.ValueOf(p)
	method := v.MethodByName("Method")
	args := []reflect.Value{reflect.ValueOf("Shenying")}
	n := method.Call(args)
	fmt.Println("Called: ", n)
}
