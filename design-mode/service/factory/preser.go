package factory

import "github.com/R0L/golang-test/design-mode/service/pres"

type Pres interface {
	Create
	Detail
	Audit
}

type Create interface {
	CreateFactory() pres.Create
}

type Detail interface {
	DetailFactory() pres.Detail
}

type Audit interface {
	AuditFactory() pres.Audit
}
