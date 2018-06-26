package decorator

import (
	"fmt"
	"strings"
)

type DispalyIntf interface {
	GetColums() int
	GetRows() int
	GetRowText(int) string
	Show()
}

type StringDisplayStu struct {
	Str string
}

func (s *StringDisplayStu) GetColums() int {
	return strings.Count(s.Str, "")
}

func (s *StringDisplayStu) GetRows() int {
	return 1
}

func (s *StringDisplayStu) GetRowText(row int) string {
	if row == 0 {
		return s.Str
	} else {
		return ""
	}
}

func (s *StringDisplayStu) Show() {
	for i := 0; i < s.GetRows(); i++ {
		fmt.Println(s.GetRowText(i))
	}
}

type SideBorderStu struct {
	BorderChar string
	Display    DispalyIntf
}

func (s *SideBorderStu) GetColums() int {
	return 1 + s.Display.GetColums() + 1
}

func (s *SideBorderStu) GetRows() int {
	return s.Display.GetRows()
}

func (s *SideBorderStu) GetRowText(row int) string {
	return s.BorderChar + s.Display.GetRowText(row) + s.BorderChar
}

func (s *SideBorderStu) Show() {
	for i := 0; i < s.GetRows(); i++ {
		fmt.Println(s.GetRowText(i))
	}
}

type FullBorderStu struct {
	Display DispalyIntf
}

func (s *FullBorderStu) GetColums() int {
	return 1 + s.Display.GetColums() + 1
}

func (s *FullBorderStu) GetRows() int {
	return 1 + s.Display.GetRows() + 1
}

func (s *FullBorderStu) GetRowText(row int) string {
	if row == 0 {
		return "+" + s.MakeLine("-", s.Display.GetColums()) + "+"
	} else if row == (s.GetRows() - 1) {
		return "+" + s.MakeLine("-", s.Display.GetColums()) + "+"
	} else {
		return "|" + s.Display.GetRowText(row-1) + "|"
	}
}

func (s *FullBorderStu) MakeLine(ch string, count int) (chs string) {
	for i := 0; i < count; i++ {
		chs += ch
	}
	return
}

func (s *FullBorderStu) Show() {
	for i := 0; i < s.GetRows(); i++ {
		fmt.Println(s.GetRowText(i))
	}
}
