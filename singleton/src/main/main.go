package main

import (
	"fmt"
	"singleton"
)

func main() {
	one := singleton.GetSingleton()
	fmt.Println(one)
}
