package product

import (
	"fmt"
)

type AppleFactory struct {
}

func (a *AppleFactory) CreatePhone() (phone, error) {
	return &applephone{}, nil
}

type applephone struct {
	Size string
}

func (b *applephone) Call() {
	fmt.Println("it's an apple mobile phone")
}
