package product

import (
	"fmt"
)

type AppleFactory struct {
}

func (a *AppleFactory) CreatePhone() (phone, error) {
	return &applephone{}, nil
}

func (a *AppleFactory) CreateComputer() (computer, error) {
	return &applecomputer{}, nil
}

type applephone struct {
	Size string
}

func (b *applephone) Call() {
	fmt.Println("it's an apple mobile phone")
}

type applecomputer struct {
	Size string
}

func (c *applecomputer) Work() {
	fmt.Println("it's an apple computer")
}
