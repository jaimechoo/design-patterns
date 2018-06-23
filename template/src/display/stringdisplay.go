package display

import (
	"fmt"
)

func NewStringDisplay(str string) displayintf {
	var one stringDisplay
	one.Str = str
	return &one
}

type stringDisplay struct {
	Str string
}

func (s *stringDisplay) Open() {
	fmt.Println("+-----+")
}

func (s *stringDisplay) Print() {
	fmt.Println("|" + s.Str + "|")
}

func (s *stringDisplay) Close() {
	fmt.Println("+-----+")
}
