package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Token = string

const (
	BALE  Token = "bale"
	KHEIR Token = "kheir"
	AREH  Token = "areh"
	NA    Token = "na"
)

type Node struct {
	Value Token
	Prev  *Node
}

func NewNode(t Token) *Node {
	return &Node{
		Value: t,
	}
}

type Stack struct {
	Head   *Node
	Length uint
}

func (s *Stack) pop() (Token, error) {
	if s.Length < 1 {
		return "", fmt.Errorf("empty stack")
	}
	temp := s.Head
	s.Head = s.Head.Prev
	temp.Prev = nil
	s.Length--
	return temp.Value, nil
}

func (s *Stack) push(t Token) {
	node := NewNode(t)
	node.Prev = s.Head
	s.Head = node
	s.Length++
}

func (s *Stack) peek() (Token, error) {
	if s.Head != nil {
		return s.Head.Value, nil
	}
	return "", fmt.Errorf("empty stack")
}

func stringAnalyst(s string) string {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'b':
			if i+3 <= len(s)-1 && string(s[i:i+4]) == BALE {
				stack.push(BALE)
				i += 3
			}
		case 'k':
			if i+4 <= len(s)-1 && string(s[i:i+5]) == KHEIR {
				if t, _ := stack.peek(); t == BALE {
					stack.pop()
					i += 4
				} else {
					return "NO"
				}
			}
		case 'a':
			if i+3 <= len(s)-1 && string(s[i:i+4]) == AREH {
				stack.push(AREH)
				i += 3
			}
		case 'n':
			if i+1 <= len(s)-1 && string(s[i:i+2]) == NA {
				if t, _ := stack.peek(); t == AREH {
					stack.pop()
					i += 1
				} else {
					return "NO"
				}
			}
		}
	}
	if _, err := stack.pop(); err == nil {
		return "NO"
	}
	return "YES"
}

var stack = &Stack{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		fmt.Println(stringAnalyst(strings.ToLower(line)))
	}
}
