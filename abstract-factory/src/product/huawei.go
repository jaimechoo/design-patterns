package product

import (
	"fmt"
)

type HuaweiFactory struct {
}

func (h *HuaweiFactory) CreatePhone() (phone, error) {
	return &huaweiphone{}, nil
}

func (x *HuaweiFactory) CreateComputer() (computer, error) {
	return &huaweicomputer{}, nil
}

type huaweiphone struct {
	Size string
}

func (c *huaweiphone) Call() {
	fmt.Println("it's a huawei mobile phone")
}

type huaweicomputer struct {
	Size string
}

func (c *huaweicomputer) Work() {
	fmt.Println("it's a huawei computer")
}
