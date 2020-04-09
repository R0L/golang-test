package reflectdemo

import "fmt"

type Demo struct {
	a string
	B int
	C interface{}
}

func NewTest(a string, b int, c interface{}) *Demo {
	return &Demo{
		a: a,
		B: b,
		C: c,
	}
}

func (t *Demo) Message() interface{} {
	return fmt.Sprintf("a:%s,b:%d,c:%v", t.a, t.B, t.C)
}
