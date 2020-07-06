package factory

import "github.com/R0L/golang-test/design-mode/service/pres"

type Doctor struct {
}

func NewDoctor() *Doctor {
	return &Doctor{}
}

func (*Doctor) CreateFactory() pres.Create {
	return pres.GetDoctorInstance()
}

func (*Doctor) DetailFactory() pres.Detail {
	return pres.GetDoctorInstance()
}

func (*Doctor) AuditFactory() pres.Audit {
	return pres.GetDoctorInstance()
}
