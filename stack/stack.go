package stack

type Stack struct {
	st  []interface{}
	len int
}

func New() *Stack {
	stack := &Stack{st: make([]interface{}, 0), len: 0}
	return stack
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Pop() {
	s.st = s.st[1:]
	s.len--
}
