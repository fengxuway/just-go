package main

import (
	"fmt"
)

type student struct {
	age int
}

type TZ int

func (t *TZ) tt (num TZ){
	*t += num
}

func main(){
	var n TZ = 100

	n.tt(40)
	fmt.Println(n)
}