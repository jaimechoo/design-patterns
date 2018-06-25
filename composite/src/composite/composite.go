package composite

import (
	"fmt"
	"strconv"
)

type EntryIntf interface {
	GetName() string
	GetSize() int
	PrintList(string)
	//	Add(entry EntryIntf)
}

func NewFileStu(name string, size int) *FileStu {
	var file FileStu
	file.Init(name, size)
	return &file
}

type FileStu struct {
	Name string
	Size int
}

func (f *FileStu) Init(name string, size int) {
	f.Name = name
	f.Size = size
}

func (f *FileStu) GetName() string {
	return f.Name
}

func (f *FileStu) GetSize() int {
	return f.Size
}

func (f *FileStu) PrintList(prefix string) {
	fmt.Printf(prefix + "/" + f.Name)
	fmt.Println("(" + strconv.Itoa(f.GetSize()) + ")")
}

//func (f *FileStu) Add(entry EntryIntf) {
//	f.Size += entry.GetSize()
//}

func NewDirectoryStu(name string) *DirectoryStu {
	var director DirectoryStu
	director.Init(name)
	return &director
}

type DirectoryStu struct {
	Name      string
	Directory []EntryIntf
}

func (d *DirectoryStu) Init(name string) {
	d.Name = name
}

func (d *DirectoryStu) GetName() string {
	return d.Name
}

func (d *DirectoryStu) GetSize() (size int) {
	for _, directory := range d.Directory {
		size += directory.GetSize()
	}
	return
}

func (d *DirectoryStu) PrintList(prefix string) {
	fmt.Printf(prefix + "/" + d.Name)
	fmt.Println("(" + strconv.Itoa(d.GetSize()) + ")")
	for _, directory := range d.Directory {
		directory.PrintList(prefix + "/" + d.Name)
	}
}

func (d *DirectoryStu) Add(entry EntryIntf) {
	d.Directory = append(d.Directory, entry)
}
