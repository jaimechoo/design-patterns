package product

import (
	"fmt"
)

type xiaomiphone struct {
	Size string
}

func (c *xiaomiphone) Call() {
	fmt.Println("it's a xiaomi mobile phone")
}
