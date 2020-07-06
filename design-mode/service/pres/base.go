package pres

import "fmt"

type base struct {
}

func NewBase() *base {
	return &base{}
}

func (*base) Create(param *CreateParam) (string, error) {
	fmt.Println("base Create")
	return "", nil
}

func (*base) Detail(transNo string) (*Info, error) {
	fmt.Println("base Detail")
	return nil, nil
}

func (*base) Audit(transNo string) (bool, error) {
	fmt.Println("base Audit")
	return false, nil
}
