package builder

import (
	"fmt"
	"os"
)

type HtmlBuilderStu struct {
	File *os.File
}

func (t *HtmlBuilderStu) MakeTitle(title string) BuilderIntf {
	filename := title + ".html"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return t
	}
	t.File = file
	t.File.Write([]byte("<html><head><title>" + title + "</title></head><body>"))
	t.File.Write([]byte("<h1>" + title + "</h1>"))
	return t
}

func (t *HtmlBuilderStu) MakeString(str string) BuilderIntf {
	if t.File == nil {
		return t
	}
	t.File.Write([]byte("<p>" + str + "</p>"))
	return t
}

func (t *HtmlBuilderStu) MakeItems(items []string) BuilderIntf {
	if t.File == nil {
		return t
	}
	t.File.Write([]byte("<ul>"))
	for _, it := range items {
		t.File.Write([]byte("<li>" + it + "</li>"))
	}
	t.File.Write([]byte("</ul>"))
	return t
}

func (t *HtmlBuilderStu) Close() {
	if t.File == nil {
		return
	}
	t.File.Write([]byte("</body></html>"))
	err := t.File.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (t *HtmlBuilderStu) GetResult() string {
	if t.File == nil {
		fmt.Println("t.File is nil")
		return ""
	}
	var tmpbytes []byte = make([]byte, 1000)
	file, err := os.Open(t.File.Name())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	length, err := file.Read(tmpbytes)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	tmpbytes = append(tmpbytes[:length], tmpbytes...)
	return string(tmpbytes)

}
