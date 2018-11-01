// Code generated by underscore_generate. DO NOT EDIT.

package stack

import (
	"fmt"
	"strings"
)

type Float32Stack interface {
	fmt.Stringer

	Push(i float32)
	Pop() float32
	Peek() float32
	Len() int
	IsEmpty() bool
	Clone() Float32Stack
}

func NewFloat32() Float32Stack {
	return &float32Stack{}
}

var _ Float32Stack = (*float32Stack)(nil)

type float32Stack struct {
	is []float32
}

func (s *float32Stack) String() string {
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

func (s *float32Stack) Push(i float32) {
	s.is = append(s.is, i)
}

func (s *float32Stack) Pop() float32 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *float32Stack) Peek() float32 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *float32Stack) Len() int {
	return len(s.is)
}

func (s *float32Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *float32Stack) Clone() Float32Stack {
	s2 := &float32Stack{is: make([]float32, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}