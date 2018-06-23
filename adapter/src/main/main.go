package main

import (
	"adapter"
)

func main() {
	var p adapter.PrintIntf = adapter.NewPrintBannerStu("Hello")
	p.PrintWeak()
	p.PrintStrong()
}
