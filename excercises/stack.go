package excercises

import (
	"../stack"
	"fmt"
)

func MakeStack() stack.Stack {
	s := stack.Stack{}
	makeIntStack := func() stack.Stack {
		for i := 1; i < 5; i++ {
			s.Push(i)
		}
		return s
	}
	return makeIntStack()
}

func MatchDelimeters(input string) bool {
	delimMap := make(map[string]string)
	delimMap["}"] = "{"
	delimMap["]"] = "["
	delimMap[")"] = "("

	isCloser := func(delim string) bool {
		_, ok := delimMap[delim]
		return ok
	}

	isOpener := func(delim string) bool {
		return delim == "{" || delim == "[" || delim == "("
	}
	// {[[[]]]}
	isMatched := func(input string) bool {
		tracker := stack.New()
		for _, char := range input {
			val := string(char)
			if isOpener(val) {
				tracker.Push(val)
			} else if isCloser(val) {
				if tracker.Len() == 0 {
					return false
				}
				last := tracker.Pop()
				if delimMap[val] != last {
					return false
				}
			}
		}
		return tracker.Len() == 0
	}

	return isMatched(input)
}

//

func stack_insertAtBottom(s stack.Stack, x interface{}) {
	if s.Len() == 0 {
		s.Push(x)
	} else {
		val := s.Peek()
		s.Pop()
		stack_insertAtBottom(s, x)
		s.Push(val)
	}
}

func stack_reverse(s stack.Stack) stack.Stack {
	if s.Len() > 0 {
		val := s.Peek()
		s.Pop()
		stack_reverse(s)
		stack_insertAtBottom(s, val)
	}
	return s
}

func ReverseStackUsingRecursion() {
	s := MakeStack()

	stack_reverse(s)

	// Print stack
	val := s.Pop()
	for val != nil {
		fmt.Print(val)
		val = s.Pop()
	}

	println()

}
