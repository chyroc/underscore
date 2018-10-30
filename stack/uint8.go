package stack

import (
	"fmt"
	"strings"
)

type Uint8Stack interface {
	fmt.Stringer

	Push(i uint8)
	Pop() uint8
	Peek() uint8
	Len() int
	IsEmpty() bool
	Clone() Uint8Stack
}

func NewUint8() Uint8Stack {
	return &uint8Stack{}
}

var _ Uint8Stack = (*uint8Stack)(nil)

type uint8Stack struct {
	is []uint8
}

func (s *uint8Stack) String() string {
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

func (s *uint8Stack) Push(i uint8) {
	s.is = append(s.is, i)
}

func (s *uint8Stack) Pop() uint8 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *uint8Stack) Peek() uint8 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *uint8Stack) Len() int {
	return len(s.is)
}

func (s *uint8Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *uint8Stack) Clone() Uint8Stack {
	s2 := &uint8Stack{is: make([]uint8, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}