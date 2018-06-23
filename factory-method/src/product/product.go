package product

type phone interface {
	Call()
}

type Phonefatory interface {
	CreatePhone() (phone, error)
}
