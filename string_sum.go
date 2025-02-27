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

const (
	errFormat = "err: %w"
)

func StringSum(input string) (output string, err error) {
	input = strings.Join(strings.Fields(input), "")
	if input == "" {
		return "", fmt.Errorf(errFormat, errorEmptyInput)
	}

	re := regexp.MustCompile(`([+-]?[^+-]+)`)
	matches := re.FindAllStringSubmatch(input, -1)
	var operands []int

	for _, v := range matches {
		intValue, err := strconv.Atoi(v[0])
		if err != nil {
			return "", fmt.Errorf(errFormat, err)
		}
		operands = append(operands, intValue)
	}

	if len(operands) != 2 {
		return "", fmt.Errorf(errFormat, errorNotTwoOperands)
	}

	return strconv.Itoa(operands[0] + operands[1]), nil
}
