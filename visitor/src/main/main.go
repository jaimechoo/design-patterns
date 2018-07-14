package main

import (
	"fmt"
	"visitor"
)

func main() {
	fmt.Println("Making root entries...")
	rootdir := visitor.NewDirectoryStu("root")
	bindir := visitor.NewDirectoryStu("bin")
	tmpdir := visitor.NewDirectoryStu("tmp")
	userdir := visitor.NewDirectoryStu("user")
	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(userdir)
	bindir.Add(visitor.NewFileStu("vi", 10000))
	bindir.Add(visitor.NewFileStu("latex", 20000))
	rootdir.Accept(new(visitor.ListVisitorStu))
}
