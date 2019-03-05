package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	val, err := Calc(str)
	if err == nil {
		fmt.Print(val)
	} else {
		fmt.Print(err)
	}
}

func Calc(str string) (int, error) {
	var stack Stack
	var result = 0
	var err error
	var isItNumber bool
	Loop:
		for _, symbol := range str {
			switch symbol {
			case ' ', '=', '\n':
				isItNumber = false
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				tempVal1 := int(symbol) - 48
				if isItNumber {
					tempVal2, tempErr := stack.Pop()
					if tempErr != nil {
						err = tempErr
						break Loop
					}
					stack.Push(tempVal2 * 10 + tempVal1)

				} else {
					isItNumber = true
					stack.Push(tempVal1)
				}
			case '+':
				isItNumber = false
				tempVal1, tempErr1 := stack.Pop()
				tempVal2, tempErr2 := stack.Pop()
				if tempErr1 != nil {
					err = tempErr1
					break Loop
				}
				if tempErr2 != nil{
					err = tempErr2
					break Loop
				}
				stack.Push(tempVal2 + tempVal1)
			case '-':
				isItNumber = false
				tempVal1, tempErr1 := stack.Pop()
				tempVal2, tempErr2 := stack.Pop()
				if tempErr1 != nil {
					err = tempErr1
					break Loop
				}
				if tempErr2 != nil{
					err = tempErr2
					break Loop
				}
				stack.Push(tempVal2 - tempVal1)
			case '*':
				isItNumber = false
				tempVal1, tempErr1 := stack.Pop()
				tempVal2, tempErr2 := stack.Pop()
				if tempErr1 != nil {
					err = tempErr1
					break Loop
				}
				if tempErr2 != nil{
					err = tempErr2
					break Loop
				}
				stack.Push(tempVal1 * tempVal2)
			case '/':
				isItNumber = false
				tempVal1, tempErr1 := stack.Pop()
				tempVal2, tempErr2 := stack.Pop()
				if tempErr1 != nil {
					err = tempErr1
					break Loop
				}
				if tempErr2 != nil{
					err = tempErr2
					break Loop
				}
				if tempVal1 == 0{
					err = errors.New("error: division by zero")
					break Loop
				}
				stack.Push(tempVal2 / tempVal1)
			case '.':
				if isItNumber{
					err = errors.New("error: float value")
				} else{
					err = errors.New("error: invalid symbol")
				}
				break Loop
			default:
				err = errors.New("error: invalid symbol")
				break Loop
			}
		}
	if err == nil{
		result, err = stack.Pop()
	}
	return result, err
}

type Stack []int

func (stack *Stack) Pop() (int, error) {
	length := len(*stack)
	if length == 0 {
		return 0, errors.New("error: stack is empty")
	}
	value := (*stack)[length-1]
	*stack = (*stack)[:length-1]
	return value, nil
}

func (stack *Stack) Push(value int) {
	*stack = append(*stack, value)
}
