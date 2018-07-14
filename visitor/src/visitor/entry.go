package visitor

import (
	"strconv"
)

type EntryIntf interface {
	Getname() string
	Getsize() int
	String() string
	Accept(VisitorIntf)
}

func NewFileStu(name string, size int) *FileStu {
	var file FileStu
	file.Name = name
	file.Size = size
	return &file
}

type FileStu struct {
	Name string
	Size int
}

func (f *FileStu) Getname() string {
	return f.Name
}

func (f *FileStu) Getsize() int {
	return f.Size
}

func (f *FileStu) String() string {
	return f.Getname() + " (" + strconv.Itoa(f.Getsize()) + ")"
}

func (f *FileStu) Accept(v VisitorIntf) {
	v.VisitFile(f)
}

func NewDirectoryStu(name string) *DirectoryStu {
	var dir DirectoryStu
	dir.Name = name
	return &dir
}

type DirectoryStu struct {
	Name      string
	Entrylist []EntryIntf
}

func (d *DirectoryStu) Getname() string {
	return d.Name
}

func (d *DirectoryStu) Getsize() (size int) {
	for _, entry := range d.Entrylist {
		size += entry.Getsize()
	}
	return
}

func (d *DirectoryStu) String() string {
	return d.Getname() + " (" + strconv.Itoa(d.Getsize()) + ")"
}

func (d *DirectoryStu) Accept(v VisitorIntf) {
	v.VisitDirectory(d)
}

func (d *DirectoryStu) Add(e EntryIntf) {
	d.Entrylist = append(d.Entrylist, e)
}
