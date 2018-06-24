package main

import (
	"fmt"
	"prototype"
)

func main() {
	//	准备
	manager := prototype.NewManagerStu()
	upen := prototype.NewUnderLinePenStu("~")
	mbox := prototype.NewMessageboxStu("*")
	sbox := prototype.NewMessageboxStu("/")
	manager.Register("strong massage", upen)
	manager.Register("warning massage", mbox)
	manager.Register("slash massage", sbox)

	//	生成
	p1 := manager.Create("strong massage")
	if p1 == nil {
		fmt.Println("p1 is nil")
	}
	p1.Use("hello world")

	p2 := manager.Create("warning massage")
	if p2 == nil {
		fmt.Println("p1 is nil")
	}
	p2.Use("hello world")

	p3 := manager.Create("slash massage")
	if p3 == nil {
		fmt.Println("p1 is nil")
	}
	p3.Use("hello world")
}
