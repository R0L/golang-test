package pres

import (
	"fmt"
	"sync"
)

type doctor struct {
	*base
}

func NewDoctor() *doctor {
	return &doctor{
		base: NewBase(),
	}
}

var once sync.Once
var instance *doctor

func GetDoctorInstance() *doctor {
	once.Do(func() {
		instance = NewDoctor()
	})
	return instance
}

func (d *doctor) Create(param *CreateParam) (string, error) {
	var (
		transNo string
		err     error
	)

	// 执行前
	fmt.Println("doctor Create before")
	if transNo, err = d.base.Create(param); nil != err {
		return transNo, err
	}
	// 执行完
	fmt.Println("doctor Create after")
	return transNo, err
}
