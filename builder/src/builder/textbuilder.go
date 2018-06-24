package builder

type TextBuilderStu struct {
	Buffer string
}

func (t *TextBuilderStu) MakeTitle(title string) BuilderIntf {
	t.Buffer += "=================================\n"
	t.Buffer += title
	t.Buffer += "\n"
	return t
}

func (t *TextBuilderStu) MakeString(str string) BuilderIntf {
	t.Buffer += " # " + str + "\n"
	t.Buffer += "\n"
	return t
}

func (t *TextBuilderStu) MakeItems(items []string) BuilderIntf {
	for _, it := range items {
		t.Buffer += " · " + it + "\n"
	}
	return t
}

func (t *TextBuilderStu) Close() {
	t.Buffer += "=================================\n"
}

func (t *TextBuilderStu) GetResult() string {
	return t.Buffer
}
