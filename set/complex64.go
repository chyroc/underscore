// Code generated by underscore_generate. DO NOT EDIT.

package set

import "sync"

type Complex64Set interface {
	Exist(i complex64) bool
	Add(i complex64)
	Adds(list ...complex64)
	Delete(i complex64)
	Deletes(list ...complex64)
	List() []complex64
}

var _ Complex64Set = (*complex64Set)(nil)

type complex64Set struct {
	m sync.Map
}

func (s *complex64Set) Exist(i complex64) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *complex64Set) Add(i complex64) {
	s.m.Store(i, struct{}{})
}

func (s *complex64Set) Adds(list ...complex64) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *complex64Set) Delete(i complex64) {
	s.m.Delete(i)
}

func (s *complex64Set) Deletes(list ...complex64) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *complex64Set) List() []complex64 {
	var list []complex64
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(complex64))
		return true
	})
	return list
}

func NewComplex64() Complex64Set {
	return &complex64Set{}
}
