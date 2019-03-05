package main

import (
	"fmt"
	"testing"
)

type successfulTestPair struct{
	number int // for comfortable counting tests and understanding of what test you thinking about now
	input string
	result int
	err string
}

var successfulTests = []successfulTestPair{
	{number: 1, input: "1 1 +", result: 2,},
	{number: 2, input: "1  1   +", result: 2,},
	{number: 3, input: "3 2 -", result: 1,},
	{number: 4, input: "2 3 *", result: 6,},
	{number: 5, input: "8 4 /", result: 2,},
	{number: 6, input: "2 4 -", result: -2,},
	{number: 7, input: "2 3 /", result: 0,},
	{number: 8, input: "2 5 * =", result: 10,},
	{number: 9, input: "4 38 4 5 6 * + - /", result: 1,},
	{number: 10, input: "100 99 - 2 * 38 + 2 3 / +", result: 40,},
}

func TestCalcSuccessful(t *testing.T) {
	for i, item := range successfulTests {
		result, err := Calc(item.input)
		if err != nil {
			fmt.Println(i)
			t.Errorf("test %d, returned message: '%v'", i + 1 , err.Error())
		}
		if result != item.result {
			t.Errorf("test %d, returned wrong result: %d, correct answer: %d", i + 1 , result ,item.result, )
		}
	}
}

var unsuccessfulTests = []successfulTestPair{
	{number: 1, input: "1 +", err: "error: stack is empty",},
	{number: 2, input: "3 0 /", err: "error: division by zero",},
	{number: 3, input: "2.3 3 !", err: "error: float value",},
	{number: 4, input: "a 4 -", err: "error: invalid symbol",},
	{number: 5, input: "2 . 4", err: "error: invalid symbol",},
	{number: 6, input: "error", err: "error: invalid symbol",},
}

func TestCalcUnsuccessful(t *testing.T) {
	for i, item := range unsuccessfulTests {
		_, err := Calc(item.input)
		if err == nil{
			t.Errorf("test %d, didn't return error, correct message: '%v'", i + 1 , item.err)
		}
		if err.Error() != item.err {
			t.Errorf("test %d, returned wrong message: '%v', correct message: '%v'", i + 1, err.Error(), item.err)
		}
	}
}