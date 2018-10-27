package stack

import (
	"fmt"
	"strings"
)

type StringStack interface {
	fmt.Stringer

	Push(i string)
	Pop() string
	Peek() string
	Len() int
	IsEmpty() bool
	Clone() StringStack
}

func NewString() StringStack {
	return &stringStack{}
}

var _ StringStack = (*stringStack)(nil)

type stringStack struct {
	is []string
}

func (s *stringStack) String() string {
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

func (s *stringStack) Push(i string) {
	s.is = append(s.is, i)
}

func (s *stringStack) Pop() string {
	if s.IsEmpty() {
		panic("is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *stringStack) Peek() string {
	if s.IsEmpty() {
		panic("is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *stringStack) Len() int {
	return len(s.is)
}

func (s *stringStack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *stringStack) Clone() StringStack {
	s2 := &stringStack{is: make([]string, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}