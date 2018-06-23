package product

import (
	"fmt"
)

type XiaomiFactory struct {
}

func (x *XiaomiFactory) CreatePhone() (phone, error) {
	return &xiaomiphone{}, nil
}

func (x *XiaomiFactory) CreateComputer() (computer, error) {
	return &xiaomicomputer{}, nil
}

type xiaomiphone struct {
	Size string
}

func (c *xiaomiphone) Call() {
	fmt.Println("it's a xiaomi mobile phone")
}

type xiaomicomputer struct {
	Size string
}

func (c *xiaomicomputer) Work() {
	fmt.Println("it's a xiaomi computer")
}
