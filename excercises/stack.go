package excercises

import (
	"../stack"
)

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
