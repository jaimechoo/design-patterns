package main

import (
	"composite"
	"fmt"
)

func main() {
	fmt.Println("Making root entries...")
	root := composite.NewDirectoryStu("root")
	bin := composite.NewDirectoryStu("bin")
	tmp := composite.NewDirectoryStu("tmp")
	user := composite.NewDirectoryStu("user")
	root.Add(bin)
	root.Add(tmp)
	root.Add(user)
	bin.Add(composite.NewFileStu("vi", 10000))
	bin.Add(composite.NewFileStu("latex", 20000))
	root.PrintList("")

	fmt.Println()
	fmt.Println("Making user entries...")
	yuki := composite.NewDirectoryStu("yuki")
	hanako := composite.NewDirectoryStu("hanako")
	tomura := composite.NewDirectoryStu("tomura")
	user.Add(yuki)
	user.Add(hanako)
	user.Add(tomura)
	yuki.Add(composite.NewFileStu("diary.html", 100))
	yuki.Add(composite.NewFileStu("Composite", 200))
	hanako.Add(composite.NewFileStu("memo.tex", 300))
	tomura.Add(composite.NewFileStu("game.doc", 400))
	tomura.Add(composite.NewFileStu("junk.mail", 500))
	root.PrintList("")
}
