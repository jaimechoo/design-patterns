package prototype

import (
	"fmt"
)

func NewUnderLinePenStu(ulchar string) *UnderLinePenStu {
	var u UnderLinePenStu
	u.Init(ulchar)
	return &u
}

type UnderLinePenStu struct {
	ulchar string
}

func (u *UnderLinePenStu) Init(ulchar string) {
	u.ulchar = ulchar
}

func (u *UnderLinePenStu) Use(str string) {
	fmt.Println(`"` + str + `"`)
	for i := 0; i < len(str); i++ {
		fmt.Print(" " + u.ulchar)
	}
	fmt.Println()
}

func (u *UnderLinePenStu) CreateClone() ProductIntf {
	var clone = &UnderLinePenStu{}
	*clone = *u
	return clone
}
