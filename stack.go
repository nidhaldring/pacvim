package main

type Stack []*NodePath

func NewStack() Stack {
	return make(Stack, 0)
}

func (s *Stack) Push(e *NodePath) {
	*s = append(*s, e)
}

func (s *Stack) Pop() *NodePath {
	l := len(*s) - 1
	elm := (*s)[l]
	*s = (*s)[:l]
	return elm
}

func (s *Stack) Contains(e *Element) bool {
	for _, elm := range *s {
		if elm.n == e {
			return true
		}
	}
	return false
}

func (s *Stack) IsNotEmpty() bool {
	return len(*s) > 0
}
