package main

import (
	"builder"
	"fmt"
)

func main() {
	var textbuilder builder.TextBuilderStu
	director := builder.NewDirectorStu(&textbuilder)
	director.ConStruct()
	fmt.Println(textbuilder.GetResult())

	var htmlbuilder builder.HtmlBuilderStu
	director2 := builder.NewDirectorStu(&htmlbuilder)
	director2.ConStruct()
	fmt.Println(htmlbuilder.GetResult())
}
