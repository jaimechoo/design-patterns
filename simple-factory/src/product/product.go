package product

import (
	"fmt"
)

const (
	Xiaomiphone int = 1 << iota
	Huaweiphone
	Applephone
)

type phone interface {
	Call()
}

func CreatPhone(Brands int) (phone, error) {
	switch Brands {
	case Xiaomiphone:
		return &xiaomiphone{}, nil
	case Huaweiphone:
		return &huaweiphone{}, nil
	case Applephone:
		return &applephone{}, nil
	default:
		return nil, fmt.Errorf("Brands not exist")
	}
}
