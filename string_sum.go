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
		return "", errorEmptyInput
	}

	re := regexp.MustCompile(`^([+-]?)(\d+)([+-])(\d+)$`)
	matches := re.FindAllStringSubmatch(input, -1)

	if len(matches) == 1 {
		firstStr := matches[0][1] + matches[0][2]
		secondStr := matches[0][3] + matches[0][4]

		firstInt, err := strconv.Atoi(firstStr)
		if err != nil {
			return "", fmt.Errorf("%s", err.Error())
		}
		secondInt, err := strconv.Atoi(secondStr)
		if err != nil {
			return "", fmt.Errorf("%s", err.Error())
		}

		return strconv.Itoa(firstInt + secondInt), nil
	}

	return "", errorNotTwoOperands
}
