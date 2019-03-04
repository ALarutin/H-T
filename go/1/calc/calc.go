package main

import (
	"errors"
	"fmt"
)

// сюда писать код
// фукция main тоже будет тут

type Stack []int

func main() {
	var s string = "3 3 +"
	val, err := Cal(s)
	if err == nil {
		fmt.Print(val)
	} else {
		fmt.Print(err)
	}
}

func Cal(str string) (int, error) {
	var stack Stack
	for _, symbol := range str {
		switch symbol {
		case '\n':
		case ' ':
		case '+':
		case '-':
		case '*':
		case '/':
		}
	}
}

func (stack *Stack) Pop() (int, error) {
	length := len(*stack)
	if length > 0 {
		value := (*stack)[length-1]
		*stack = (*stack)[:length-1]
		return value, nil
	}
	return 0, errors.New("Stack is empty")
}

func (stack *Stack) Push(value int) {
	*stack = append(*stack, value)
}
