package product

import (
	"fmt"
)

type HuaweiFactory struct {
}

func (h *HuaweiFactory) CreatePhone() (phone, error) {
	return &huaweiphone{}, nil
}

type huaweiphone struct {
	Size string
}

func (c *huaweiphone) Call() {
	fmt.Println("it's a huawei mobile phone")
}
