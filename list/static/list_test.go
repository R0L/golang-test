package static_test

import (
	"github.com/R0L/golang-test/list/static"
	"os"
	"testing"
)

var l *static.List

func TestMain(m *testing.M) {
	l = static.Init()
	os.Exit(m.Run())
}

// 测试链表是否为空
func TestEmpty(t *testing.T) {
	TestInsert(t)
	flag, err := static.Empty(l)
	t.Log(static.ToArray(l), flag, err)
}

// 测试链表清空
func TestClear(t *testing.T) {
	TestInsert(t)
	err := static.Clear(l)
	t.Log(static.ToArray(l), err)
}

// 测试链表获取
func TestGet(t *testing.T) {
	TestInsert(t)
	var listIndex = []int{
		-1, 1, 2, 3, 12, 17, 19, 20, 21,
	}
	for _, index := range listIndex {
		val, err := static.Get(l, index)
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
		index, err := static.Locate(l, element)
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
		err := static.Insert(l, item.Key, item.Value)
		t.Log(static.ToArray(l), static.Length(l), err)
	}
}

// 测试链表删除
func TestDelete(t *testing.T) {
	TestInsert(t)
	var listIndex = []int{
		20, 21, 20, 20, 20, 20, 1, 2, 3, 4, 7, 8, 9, 12, 13,
	}
	for _, index := range listIndex {
		val, err := static.Delete(l, index)
		t.Log(index, val, static.Length(l), static.ToArray(l), err)
	}
}

// 测试链表长度
func TestLength(t *testing.T) {
	TestInsert(t)
	length := static.Length(l)
	t.Log(length)
}
