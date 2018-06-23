package main

import (
	"display"
)

func main() {
	one := display.NewDispaly(display.NewCharDisplay("H"))
	one.Display()

	two := display.NewDispaly(display.NewStringDisplay("hello world"))
	two.Display()
}
