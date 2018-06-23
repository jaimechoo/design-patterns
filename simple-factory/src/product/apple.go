package product

import (
	"fmt"
)

type applephone struct {
	Size string
}

func (b *applephone) Call() {
	fmt.Println("it's an apple mobile phone")
}
