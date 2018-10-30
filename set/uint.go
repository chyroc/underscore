package set

import "sync"

type UintSet interface {
	Exist(i uint) bool
	Add(i uint)
	Adds(list ...uint)
	Delete(i uint)
	Deletes(list ...uint)
	List() []uint
}

var _ UintSet = (*uintSet)(nil)

type uintSet struct {
	m sync.Map
}

func (s *uintSet) Exist(i uint) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *uintSet) Add(i uint) {
	s.m.Store(i, struct{}{})
}

func (s *uintSet) Adds(list ...uint) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *uintSet) Delete(i uint) {
	s.m.Delete(i)
}

func (s *uintSet) Deletes(list ...uint) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *uintSet) List() []uint {
	var list []uint
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(uint))
		return true
	})
	return list
}

func NewUint() UintSet {
	return &uintSet{}
}
