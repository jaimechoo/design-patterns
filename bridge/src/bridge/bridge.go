package bridge

import (
	"fmt"
)

func NewDisPlayStu(impl DisplayImplIntf) *DisplayStu {
	var display DisplayStu
	display.Init(impl)
	return &display
}

type DisplayStu struct {
	impl DisplayImplIntf
}

func (d *DisplayStu) Init(impl DisplayImplIntf) {
	d.impl = impl
}

func (d *DisplayStu) Open() {
	d.impl.RawOpen()
}

func (d *DisplayStu) Print() {
	d.impl.RawPrint()
}

func (d *DisplayStu) Close() {
	d.impl.RawClose()
}

func (d *DisplayStu) DisPlay() {
	d.Open()
	d.Print()
	d.Close()
}

func NewCountDisplayStu(impl DisplayImplIntf) *CountDisplayStu {
	var display CountDisplayStu
	display.DisplayStu.Init(impl)
	return &display
}

type CountDisplayStu struct {
	DisplayStu
}

func (c *CountDisplayStu) MultiDisplay(times int) {
	c.Open()
	for i := 0; i < times; i++ {
		c.Print()
	}
	c.Close()
}

type DisplayImplIntf interface {
	RawOpen()
	RawPrint()
	RawClose()
}

type StringDisplayImplStu struct {
	Str string
}

func (s *StringDisplayImplStu) RawOpen() {
	s.Printline()
}

func (s *StringDisplayImplStu) RawPrint() {
	fmt.Println("|" + s.Str + "|")
}

func (s *StringDisplayImplStu) RawClose() {
	s.Printline()
}

func (s *StringDisplayImplStu) Printline() {
	fmt.Print("+")
	for i := 0; i < len(s.Str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
