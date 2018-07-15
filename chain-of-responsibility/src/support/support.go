package support

import (
	"fmt"
)

type SupportIntf interface {
	Support(TroubleStu)
	SetNext(SupportIntf) SupportIntf
	Done(TroubleStu)
	Fail(TroubleStu)
	Resovle(TroubleStu) bool
}

type SupportStu struct {
	Name string
	This SupportIntf
	Next SupportIntf
}

func (s *SupportStu) String() string {
	return "[" + s.Name + "]"
}

func (s *SupportStu) SetNext(ss SupportIntf) SupportIntf {
	s.Next = ss
	return ss
}

func (s *SupportStu) Support(t TroubleStu) {
	if s.This.Resovle(t) {
		s.Done(t)
	} else if s.Next != nil {
		s.Next.Support(t)
	} else {
		s.Fail(t)
	}
}

func (s *SupportStu) Done(t TroubleStu) {
	fmt.Println(t.String() + " is resolved by " + s.String() + ".")
}

func (s *SupportStu) Fail(t TroubleStu) {
	fmt.Println(t.String() + " cannot be resolved.")
}

func NewNoSupportStu(name string) *NoSupportStu {
	var n NoSupportStu
	n.Name = name
	n.This = &n
	return &n
}

type NoSupportStu struct {
	SupportStu
}

func (n *NoSupportStu) Resovle(t TroubleStu) bool {
	return false
}

func NewLimitSupportStu(name string, limit int) *LimitSupportStu {
	var l LimitSupportStu
	l.Name = name
	l.Limit = limit
	l.This = &l
	return &l
}

type LimitSupportStu struct {
	Limit int
	SupportStu
}

func (l *LimitSupportStu) Resovle(t TroubleStu) bool {
	if t.GetNumber() < l.Limit {
		return true
	} else {
		return false
	}
}

func NewOddSupportStu(name string) *OddSupportStu {
	var o OddSupportStu
	o.Name = name
	o.This = &o
	return &o
}

type OddSupportStu struct {
	SupportStu
}

func (o *OddSupportStu) Resovle(t TroubleStu) bool {
	if t.GetNumber()%2 == 1 {
		return true
	} else {
		return false
	}
}

func NewSpecialSupportStu(name string, number int) *SpecialSupportStu {
	var s SpecialSupportStu
	s.Name = name
	s.Number = number
	s.This = &s
	return &s
}

type SpecialSupportStu struct {
	Number int
	SupportStu
}

func (s *SpecialSupportStu) Resovle(t TroubleStu) bool {
	if t.GetNumber() == s.Number {
		return true
	} else {
		return false
	}
}
