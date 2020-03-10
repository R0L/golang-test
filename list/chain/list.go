package chain

import "C"
import (
	"errors"
)

const (
	ErrOutOfChainList       = "out of chain list"
	ErrChainListAlreadyFull = "chain list already full"
	ErrChainListIsEmpty     = "chain list is nil"
)

//线性表的数据对象集合为(a1,a2,...,a3)，每个元素的类型均为DataType。其中，除第一个元素a1外，每一个元素有且只有一个直接前驱元素，除了最后一个元素an外，每一个元素有且只有一个直接后继元素。数据元素之间的关系是一对一的关系。
// 节点
type Node struct {
	Value interface{}
	Node  *Node
}
type List struct {
	*Node
	Length int
}

// 初始化操作，建立一个空的线性表L。
func Init() *List {
	return &List{}
}

// 若线性表为空，返回true，否则返回false。
func Empty(list *List) (bool, error) {
	if nil == list {
		return false, errors.New(ErrChainListIsEmpty)
	}

	if list.Length == 0 {
		return true, nil
	}

	return false, nil
}

// 将线性表清空。
func Clear(list *List) error {
	if nil == list {
		return errors.New(ErrChainListIsEmpty)
	}
	list.Length = 0
	return clearNode(list.Node)
}
func clearNode(node *Node) error {
	if nil == node {
		return nil
	}
	if nil == node.Value {
		return nil
	}
	newNode := node.Node
	return clearNode(newNode)
}

// 将线性表L中的第i个位置元素值返回给e。
func Get(list *List, i int) (interface{}, error) {
	if nil == list {
		return nil, errors.New(ErrChainListIsEmpty)
	}
	if i <= 0 || i > list.Length {
		return nil, errors.New(ErrOutOfChainList)
	}

	newNode := list.Node
	if nil == newNode {
		return nil, nil
	}
	for ; i > 0; i-- {
		newNode = newNode.Node
	}

	return newNode.Value, nil
}

// 在线性表L中查找与给定值e相等的元素，如果查找成功，返回该元素在比表中序号表示成功；否则，返回0表示失败。
func Locate(list *List, e interface{}) (int, error) {
	var findKey int
	if nil == list {
		return findKey, errors.New(ErrChainListIsEmpty)
	}
	newNode := list.Node
	if nil == newNode {
		return findKey, nil
	}
	for i := 0; i < list.Length; i++ {
		if e == newNode.Value {
			findKey = i
			break
		}
		newNode = newNode.Node
	}

	return findKey, nil
}

// 在线性表L中的第i个位置插入新元素e。
func Insert(list *List, i int, e interface{}) error {
	if nil == list {
		return errors.New(ErrChainListIsEmpty)
	}

	if i <= 0 || i > list.Length+1 {
		return errors.New(ErrOutOfChainList)
	}

	newNode := list.Node
	if nil == newNode {
		return nil
	}
	for ; i > 1; i-- {
		newNode = newNode.Node
	}

	newNode.Node = &Node{
		Value: e,
		Node:  newNode.Node,
	}
	list.Length += 1

	return nil
}

// 删除线性表L中第i个位置元素，并用e返回其值。
func Delete(list *List, i int) (interface{}, error) {
	var retVal interface{}
	if nil == list {
		return nil, errors.New(ErrChainListIsEmpty)
	}
	if 0 >= i || i > list.Length {
		return nil, errors.New(ErrOutOfChainList)
	}

	newNode := list.Node
	if nil == newNode {
		return nil, nil
	}
	for ; i > 1; i-- {
		newNode = newNode.Node
	}

	newNode.Node = newNode.Node.Node
	list.Length -= 1

	return retVal, nil
}

// 返回线性表L的元素个数。
func Length(list *List) int {
	if nil == list {
		return 0
	}
	return list.Length
}

// 遍历
func ToArray(list *List) []interface{} {
	var ret []interface{}
	newNode := list.Node
	for i := list.Length; i > 0; i-- {
		newNode = newNode.Node
		if nil == newNode.Value {
			break
		}
		ret = append(ret, newNode.Value)
	}
	return ret
}
