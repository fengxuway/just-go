package main

import (
	"fmt"
)

type USB interface {
	Name() string
	connect
}

type connect interface {
	Connect()
}

type USBJ struct {
	name string
}
type USBK struct {
	name string
}

func (usbj USBJ) Name() string {
	return usbj.name
}

func (usbj USBJ) Connect() {
	fmt.Println("USBJ connected.")
}

func (usbj USBK) Name() string {
	return usbj.name
}

func (usbj USBK) Connect() {
	fmt.Println("USBJ connected.")
}

func main(){
	var usb USB
	usb = USBJ{"usb instance"}
	CCC(usb)
}

func CCC(usb USB){
	if p, ok := usb.(USBJ); ok{
		// 判断是否为某个类型，第一个返回值p为强制转换后的struct
		fmt.Println("usb is USBJ, name: ", p.name)
	}
	// var.(type)常用于switch判断变量的类型
	switch p:=usb.(type){
	case USBJ:
		fmt.Println(p.name)
	}
	fmt.Println("Start connect : ", usb.Name())
	usb.Connect()
}