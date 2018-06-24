package prototype

import (
	"fmt"
)

func NewMessageboxStu(decochar string) *MessageboxStu {
	var m MessageboxStu
	m.Init(decochar)
	return &m
}

type MessageboxStu struct {
	decochar string
}

func (m *MessageboxStu) Init(decochar string) {
	m.decochar = decochar
}

func (m *MessageboxStu) Use(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Print(m.decochar)
	}
	fmt.Println()
	fmt.Println(m.decochar + " " + str + " " + m.decochar)
	for i := 0; i < len(str); i++ {
		fmt.Print(m.decochar)
	}
	fmt.Println()
}

func (m *MessageboxStu) CreateClone() ProductIntf {
	var clone = &MessageboxStu{}
	*clone = *m
	return clone
}
