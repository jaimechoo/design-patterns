package visitor

import (
	"fmt"
)

type VisitorIntf interface {
	VisitFile(EntryIntf)
	VisitDirectory(EntryIntf)
}

type ListVisitorStu struct {
	Currentdir string
}

func (l *ListVisitorStu) VisitFile(file EntryIntf) {
	fmt.Println(l.Currentdir + "/" + file.String())
}

func (l *ListVisitorStu) VisitDirectory(dir EntryIntf) {
	fmt.Println(l.Currentdir + "/" + dir.String())
	savedir := l.Currentdir
	l.Currentdir = l.Currentdir + "/" + dir.Getname()
	for _, entry := range dir.(*DirectoryStu).Entrylist {
		entry.Accept(l)
	}
	l.Currentdir = savedir
}
