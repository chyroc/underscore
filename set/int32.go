package set

import "sync"

type Int32Set interface {
	Exist(i int32) bool
	Add(i int32)
	Adds(list ...int32)
	Delete(i int32)
	Deletes(list ...int32)
	List() []int32
}

var _ Int32Set = (*int32Set)(nil)

type int32Set struct {
	m sync.Map
}

func (s *int32Set) Exist(i int32) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *int32Set) Add(i int32) {
	s.m.Store(i, struct{}{})
}

func (s *int32Set) Adds(list ...int32) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *int32Set) Delete(i int32) {
	s.m.Delete(i)
}

func (s *int32Set) Deletes(list ...int32) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *int32Set) List() []int32 {
	var list []int32
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(int32))
		return true
	})
	return list
}

func NewInt32() Int32Set {
	return &int32Set{}
}
