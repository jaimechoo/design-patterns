package display

import (
	"fmt"
)

func NewCharDisplay(ch string) displayintf {
	var one charDisplay
	one.Ch = ch
	return &one
}

type charDisplay struct {
	Ch string
}

func (c *charDisplay) Open() {
	fmt.Print("<<")
}

func (c *charDisplay) Print() {
	fmt.Print(c.Ch)
}

func (c *charDisplay) Close() {
	fmt.Println(">>")
}
