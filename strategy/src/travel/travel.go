package travel

type vehiclelIntf interface {
	Move(from, to string)
}

type Vehicle struct {
	Departure   string
	Destination string
	Vehiclel    vehiclelIntf
}

func (o *Vehicle) Init(departure, destination string, vehiclel vehiclelIntf) {
	o.Departure = departure
	o.Destination = destination
	o.Vehiclel = vehiclel
}

func (o *Vehicle) Move() {
	o.Vehiclel.Move(o.Departure, o.Destination)
}
