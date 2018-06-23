package product

type phone interface {
	Call()
}

type computer interface {
	Work()
}

type Electronicfatory interface {
	CreatePhone() (phone, error)
	CreateComputer() (computer, error)
}
