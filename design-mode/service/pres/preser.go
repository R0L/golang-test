package pres

type CreateParam struct {
}

type Create interface {
	Create(param *CreateParam) (string, error)
}

type Info struct {
}

type Detail interface {
	Detail(transNo string) (*Info, error)
}

type Audit interface {
	Audit(transNo string) (bool, error)
}
