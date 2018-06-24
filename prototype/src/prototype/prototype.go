package prototype

type ProductIntf interface {
	Use(string)
	CreateClone() ProductIntf
}

func NewManagerStu() *ManagerStu {
	var m ManagerStu
	m.Init()
	return &m
}

type ManagerStu struct {
	showcase map[string]ProductIntf
}

func (m *ManagerStu) Init() {
	m.showcase = make(map[string]ProductIntf)
}

func (m *ManagerStu) Register(name string, product ProductIntf) {
	m.showcase[name] = product
}

func (m *ManagerStu) Create(name string) ProductIntf {
	product, ok := m.showcase[name]
	if ok {
		return product.CreateClone()
	} else {
		return nil
	}
}
