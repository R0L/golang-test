package pres_test

import (
	"github.com/R0L/golang-test/design-mode/controller/pres"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	pres.Create()
}

func TestDetail(t *testing.T) {
	pres.Detail()
}

func TestAudit(t *testing.T) {
	pres.Audit()
}
