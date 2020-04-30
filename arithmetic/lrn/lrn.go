package lrn

import "container/list"

type LRN struct {
	len  int
	list *list.List
}

func NewLRN(len int) *LRN {
	return &LRN{
		len:  len,
		list: list.New(),
	}
}

// 检查
func (s *LRN) check(key string) (*list.Element, error) {
	for ele := s.list.Back(); ele != nil; {
		if ele.Value.(string) == key {
			return ele, nil
		}
		ele = ele.Next()
	}
	return nil, nil
}

// 添加
func (s *LRN) add(key string) error {
	if s.list.Len() >= s.len {
		s.list.Remove(s.list.Back())
	}
	s.list.PushFront(key)
	return nil
}

// 移动
func (s *LRN) move(ele *list.Element) error {
	s.list.MoveToFront(ele)
	return nil
}

// lrn
func (s *LRN) Lrn(keys ...string) error {
	for _, key := range keys {
		if err := s.singe(key); nil != err {
			return err
		}
	}
	return nil
}

func (s *LRN) singe(key string) error {
	ele, err := s.check(key)
	if nil != err {
		return err
	}
	if ele == nil {
		return s.add(key)
	}
	return s.move(ele)
}
