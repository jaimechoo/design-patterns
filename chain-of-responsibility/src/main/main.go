package main

import (
	"support"
)

func main() {
	alice := support.NewNoSupportStu("Alice")
	bob := support.NewLimitSupportStu("Bob", 100)
	charlie := support.NewSpecialSupportStu("Charlie", 429)
	diana := support.NewLimitSupportStu("Diana", 200)
	elmo := support.NewOddSupportStu("Elmo")
	fred := support.NewLimitSupportStu("Fred", 300)
	//形成职责链
	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)
	//制造各种问题
	for i := 0; i < 500; i += 33 {
		alice.Support(support.TroubleStu{Number: i})
	}
}
