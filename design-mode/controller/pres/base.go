package pres

import (
	"github.com/R0L/golang-test/design-mode/service/factory"
	"github.com/R0L/golang-test/design-mode/service/pres"
)

func create(creater factory.Create, param *pres.CreateParam) (string, error) {
	return creater.CreateFactory().Create(param)
}

func detail(detailer factory.Detail, transNo string) (*pres.Info, error) {
	return detailer.DetailFactory().Detail(transNo)
}

func audit(auditer factory.Audit, transNo string) (bool, error) {
	return auditer.AuditFactory().Audit(transNo)
}
