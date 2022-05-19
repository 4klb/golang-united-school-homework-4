package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5h ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var str string
	var whitespace int
	var firstNeg bool
	var sum int

	for _, value := range input {
		if value == ' ' {
			whitespace++
			continue
		}
		str += string(value)
	}

	if (whitespace != 0 && whitespace == len(input)) || (len(input) == 0) {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	s := strings.Split(str, " ")

	if s[0][0] == '-' {
		firstNeg = true
	}

	if firstNeg {
		str = s[0][1:]
	}

	res, symbol := SplitBySymbol(str)

	if len(res) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	num1, num2, err := StrConvToInt(res)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if firstNeg && symbol == '+' {
		sum = num2 - num1
	} else if firstNeg && symbol == '-' {
		sum = num2 + num1
		sum *= -1
	} else if !firstNeg && symbol == '+' {
		sum = num1 + num2
	} else if num1 < num2 {
		sum = num1 - num2
	}

	strResult := IntConvToStr(sum)

	return strResult, nil
}

func SplitBySymbol(str string) ([]string, rune) {
	var res []string
	var symbol rune

	for _, value := range str {
		if value == '-' {
			res = strings.Split(str, "-")
			symbol = FindSymbol(value)
			break
		} else if value == '+' {
			res = strings.Split(str, "+")
			symbol = FindSymbol(value)
			break
		}
	}
	return res, symbol
}

func FindSymbol(symbol rune) rune {
	if symbol == '-' {
		return '-'
	} else if symbol == '+' {
		return '+'
	}
	return ' '
}

func IntConvToStr(sum int) string {
	res := strconv.Itoa(sum)
	return res
}

func StrConvToInt(arr []string) (int, int, error) {
	num1, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, err
	}

	num2, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}
