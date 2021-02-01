package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	parenthesizedFormat = "\\( ([0-9]+) ([\\+\\/\\*-]) ([0-9]+) \\)"
)

// infixCalculator recursively reduces parentheses via regex
func infixCalculator(s string) (int, error) {
	trimSuffix := strings.TrimSuffix(s, " ")
	// Base case is just the integer left
	val, isInt := isInteger(trimSuffix)
	if isInt {
		return val, nil
	}
	parenRegex, err := regexp.Compile(parenthesizedFormat)
	if err != nil {
		return -1, err
	}
	// Find the substrings that match the regex (and can be reduced) and their indices.
	matches := parenRegex.FindAllStringSubmatch(s, -1)
	indices := parenRegex.FindAllStringSubmatchIndex(s, -1)

	if len(matches) == 0 {
		return -1, errors.New("invalid format, found no matches")
	}

	calculations := []int{}
	// Go through matches and calculate the sum in the brackets
	for _, match := range matches {
		num1, _ := isPositiveInteger(match[1])
		operator := match[2]
		num2, _ := isPositiveInteger(match[3])
		val := performCalculation(operator, num1, num2)
		calculations = append(calculations, val)
	}

	// Rebuild the new string for next reduction phase
	newString := ""
	i := 0
	for x, index := range indices {
		start := index[0]
		end := index[1]
		newString += s[i:start] + strconv.Itoa(calculations[x])
		i = end
	}
	newString += s[i:]
	return infixCalculator(newString)
}