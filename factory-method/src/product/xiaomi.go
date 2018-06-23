package product

import (
	"fmt"
)

type XiaomiFactory struct {
}

func (x *XiaomiFactory) CreatePhone() (phone, error) {
	return &xiaomiphone{}, nil
}

type xiaomiphone struct {
	Size string
}

func (c *xiaomiphone) Call() {
	fmt.Println("it's a xiaomi mobile phone")
}
