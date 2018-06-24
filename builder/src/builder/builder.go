package builder

type BuilderIntf interface {
	MakeTitle(string) BuilderIntf
	MakeString(string) BuilderIntf
	MakeItems([]string) BuilderIntf
	Close()
}

func NewDirectorStu(builder BuilderIntf) *DirectorStu {
	var director DirectorStu
	director.Init(builder)
	return &director
}

type DirectorStu struct {
	builder BuilderIntf
}

func (d *DirectorStu) Init(builder BuilderIntf) {
	d.builder = builder
}

func (d *DirectorStu) ConStruct() {
	d.builder.MakeTitle("Greating")
	d.builder.MakeString("从早上至下午")
	d.builder.MakeItems([]string{
		"早上好",
		"下午好",
	})
	d.builder.MakeString("晚上")
	d.builder.MakeItems([]string{
		"晚上好",
		"晚安",
		"再见",
	})
	d.builder.Close()
}
