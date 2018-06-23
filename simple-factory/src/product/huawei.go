package product

import (
	"fmt"
)

type huaweiphone struct {
	Size string
}

func (c *huaweiphone) Call() {
	fmt.Println("it's a huawei mobile phone")
}
