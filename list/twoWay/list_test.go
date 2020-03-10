package twoWay_test

import (
	"github.com/R0L/golang-test/list/twoWay"
	"os"
	"testing"
)

var l *twoWay.List

func TestMain(m *testing.M) {
	l = twoWay.Init()
	os.Exit(m.Run())
}

// 测试链表是否为空
func TestEmpty(t *testing.T) {
	TestInsert(t)
	flag, err := twoWay.Empty(l)
	t.Log(twoWay.ToArray(l), flag, err)
}

// 测试链表清空
func TestClear(t *testing.T) {
	TestInsert(t)
	err := twoWay.Clear(l)
	t.Log(twoWay.ToArray(l), err)
}

// 测试链表获取
func TestGet(t *testing.T) {
	TestInsert(t)
	var listIndex = []int{
		-1, 1, 2, 3, 19, 20, 21,
	}
	for _, index := range listIndex {
		val, err := twoWay.Get(l, index)
		t.Log(index, val, err)
	}

}

// 测试链表查询
func TestLocate(t *testing.T) {
	TestInsert(t)
	var listElement = []int{
		-1, 2,
	}
	for _, element := range listElement {
		index, err := twoWay.Locate(l, element)
		t.Log(element, index, err)
	}
}

// 测试链表查询
func TestInsert(t *testing.T) {
	var params = [...]struct {
		Key   int
		Value interface{}
	}{
		// 正常测试
		{1, 1},
		{1, 2},
		{1, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{2, 2},
		// 异常测试
		{0, 0},
		{-1, -1},
		{9, 9},
		{10, 10},

		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
	}
	for _, item := range params {
		t.Log(item.Key, item.Value)
		err := twoWay.Insert(l, item.Key, item.Value)
		t.Log(twoWay.ToArray(l), err)
	}
}

// 测试链表删除
func TestDelete(t *testing.T) {
	TestInsert(t)
	var listIndex = []int{
		20, 21, 20, 20, 20, 20, 1, 2, 3, 4,
	}
	for _, index := range listIndex {
		val, err := twoWay.Delete(l, index)
		t.Log(index, val, l.Length, twoWay.ToArray(l), err)
	}
}

// 测试链表长度
func TestLength(t *testing.T) {
	TestInsert(t)
	length := twoWay.Length(l)
	t.Log(length)
}
