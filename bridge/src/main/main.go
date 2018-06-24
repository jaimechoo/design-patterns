package main

import (
	"bridge"
)

func main() {
	d1 := bridge.NewDisPlayStu(&bridge.StringDisplayImplStu{Str: "Hello,China"})
	d2 := bridge.NewCountDisplayStu(&bridge.StringDisplayImplStu{Str: "Hello,World"})
	d1.DisPlay()
	d2.DisPlay()
	d2.MultiDisplay(5)
}
