package reflectdemo_test

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type Content struct {
	sync.Once
}

func (c *Content) NewA() *A {
	var a *A
	c.Once.Do(func() {
		a = &A{}
	})
	return a
}

type A struct {
}

func (*A) Test() {
	fmt.Println("ab")
}

func TestReflect(t *testing.T) {
	//d := &reflectdemo.Demo{}
	//rv := reflect.ValueOf(d)
	//fmt.Println(rv, rv.Kind(), rv.Elem(), rv.Elem().Kind())
	//rt := reflect.TypeOf(d)
	//fmt.Println(rt, rt.Kind(), rt.Elem(), rt.Elem().Kind())

	//var a = 1
	//
	//rv := reflect.ValueOf(a)
	//
	//if rv.Kind() == reflect.Ptr {
	//	rv = rv.Elem()
	//}
	//if rv.CanSet() {
	//	switch rv.Kind() {
	//	case reflect.Int:
	//		rv.SetInt(12)
	//
	//	}
	//}
	//
	//
	//fmt.Println(a)

	structName := "A"
	funcName := "Test"

	fName := fmt.Sprintf("New%s", structName)
	//fmt.Println(reflect.ValueOf(NewA))
	//fmt.Println(reflect.ValueOf("NewA".(interface{})))
	rs := reflect.ValueOf(&Content{}).MethodByName(fName).Call(nil)
	rs[0].MethodByName(funcName).Call(nil)

}
