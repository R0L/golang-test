package sequence

import "errors"

const (
	DefaultMaxSize          = 20
	ErrOutOfChainList       = "out of chain list"
	ErrChainListAlreadyFull = "chain list already full"
	ErrChainListIsEmpty     = "chain list is nil"
)

//线性表的数据对象集合为(a1,a2,...,a3)，每个元素的类型均为DataType。其中，除第一个元素a1外，每一个元素有且只有一个直接前驱元素，除了最后一个元素an外，每一个元素有且只有一个直接后继元素。数据元素之间的关系是一对一的关系。
type List struct {
	InnerList [DefaultMaxSize]interface{}
	MaxSize   int
	Length    int
}

// 初始化操作，建立一个空的线性表L。
func Init() *List {
	return &List{
		InnerList: [DefaultMaxSize]interface{}{},
		MaxSize:   DefaultMaxSize,
		Length:    0,
	}
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
	list.InnerList = [DefaultMaxSize]interface{}{}
	list.Length = 0
	return nil
}

// 将线性表L中的第i个位置元素值返回给e。
func Get(list *List, i int) (interface{}, error) {
	if nil == list {
		return nil, errors.New(ErrChainListIsEmpty)
	}
	if i <= 0 || i > list.Length {
		return nil, errors.New(ErrOutOfChainList)
	}

	return list.InnerList[i-1], nil
}

// 在线性表L中查找与给定值e相等的元素，如果查找成功，返回该元素在比表中序号表示成功；否则，返回0表示失败。
func Locate(list *List, e interface{}) (int, error) {
	var findKey int
	if nil == list {
		return findKey, errors.New(ErrChainListIsEmpty)
	}
	for key, value := range list.InnerList {
		if e == value {
			findKey = key
			break
		}
	}

	return findKey, nil
}

// 在线性表L中的第i个位置插入新元素e。
func Insert(list *List, i int, e interface{}) error {
	if nil == list {
		return errors.New(ErrChainListIsEmpty)
	}
	if list.Length >= list.MaxSize {
		return errors.New(ErrChainListAlreadyFull)
	}

	if i <= 0 || i > list.Length+1 {
		return errors.New(ErrOutOfChainList)
	}

	for j := list.Length; j > i-1; j-- {
		list.InnerList[j] = list.InnerList[j-1]
	}
	list.InnerList[i-1] = e
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

	retVal = list.InnerList[i-1]

	for j := i - 1; j < list.Length-1; j++ {
		list.InnerList[j] = list.InnerList[j+1]
	}
	list.InnerList[list.Length-1] = nil
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
