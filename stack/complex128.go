// Code generated by underscore_generate. DO NOT EDIT.

package stack

import (
	"fmt"
	"strings"
)

type Complex128Stack interface {
	fmt.Stringer

	Push(i complex128)
	Pop() complex128
	Peek() complex128
	Len() int
	IsEmpty() bool
	Clone() Complex128Stack
}

func NewComplex128() Complex128Stack {
	return &complex128Stack{}
}

var _ Complex128Stack = (*complex128Stack)(nil)

type complex128Stack struct {
	is []complex128
}

func (s *complex128Stack) String() string {
	var buf = new(strings.Builder)
	for i, v := range s.is {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("%v", v))
		} else {
			buf.WriteString(fmt.Sprintf(" < %v", v))
		}
	}
	return buf.String()
}

func (s *complex128Stack) Push(i complex128) {
	s.is = append(s.is, i)
}

func (s *complex128Stack) Pop() complex128 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *complex128Stack) Peek() complex128 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *complex128Stack) Len() int {
	return len(s.is)
}

func (s *complex128Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *complex128Stack) Clone() Complex128Stack {
	s2 := &complex128Stack{is: make([]complex128, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}