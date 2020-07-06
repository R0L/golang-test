package pres

import (
	"fmt"
	"github.com/R0L/golang-test/design-mode/service/factory"
	"github.com/R0L/golang-test/design-mode/service/pres"
)

var doctorFactory *factory.Doctor

func init() {
	doctorFactory = factory.NewDoctor()
}

func Create() {

	param := &pres.CreateParam{}
	transNo, err := create(doctorFactory, param)

	fmt.Printf("pres Create transNo:%v, err:%v\n", transNo, err)

}

func Detail() {

	transNo := ""
	info, err := detail(doctorFactory, transNo)

	fmt.Printf("pres Create info:%v, err:%v", info, err)

}

func Audit() {

	transNo := ""
	flag, err := audit(doctorFactory, transNo)

	fmt.Printf("pres Create flag:%v, err:%v", flag, err)

}
