package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.Join(strings.Fields(input), "")
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	var hasSymboleReg = regexp.MustCompile(`([+-]?[^+-]+)`)
	symbols := hasSymboleReg.FindAllStringSubmatch(input, -1)
	var numb []int

	for _, v := range symbols {
		value, err := strconv.Atoi(v[0])
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}
		numb = append(numb, value)
	}

	if len(numb) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	result := strconv.Itoa(numb[0] + numb[1])
	fmt.Println(result)
	return result, nil
}
