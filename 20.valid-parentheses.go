package main

import (
	"fmt"
)

func main() {
	input := "{[]}"
	output := isValid(input)
	fmt.Println(output)
}


/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	ml := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	mr := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := New()

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := ml[c]; ok {
			stack.Push(c)
		} else {
			if l, ok := mr[c]; ok {
				if stack.Peek() == l {
					stack.Pop()
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return stack.Len() == 0
}

type (
	// Stack is a linear data structure which follows a particular order in which the operations are performed.
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

//https://github.com/golang-collections/collections/blob/master/stack/stack.go

// New Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Len Return the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// Peek View the top item on the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop Pops the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push Push a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

// @lc code=end

