package display

type displayintf interface {
	Open()
	Print()
	Close()
}

func NewDispaly(in displayintf) *dispaly {
	var one dispaly
	one.intf = in
	return &one
}

type dispaly struct {
	intf displayintf
}

func (d *dispaly) Display() {
	d.intf.Open()
	for i := 0; i < 5; i++ {
		d.intf.Print()
	}
	d.intf.Close()
}
