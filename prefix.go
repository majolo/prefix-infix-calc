package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// prefixCalculator computes a prefix notation input and returns the value.
// If an error is returned the result is -1.
func prefixCalculator(s string) (result int, err error) {
	values := strings.Split(s, " ")
	nodes := []Node{}
	for _, value := range values {
		num, isInt := isPositiveInteger(value)
		isOp := isOperator(value)
		if !(isInt || isOp) {
			return -1, errors.New(fmt.Sprintf("input string contained invalid characters %s", value))
		}
		if isInt {
			nodes = append(nodes, CalculatorInt{value: num})
			continue
		}
		if isOp {
			nodes = append(nodes, CalculatorOperator{operator: value})
			continue
		}
	}
	return reduceNodes(nodes)
}

// reduceNodes takes the tokenized input and recursively reduces from the right hand side.
func reduceNodes(nodes []Node) (result int, err error) {
	// Handle base cases
	if len(nodes) == 0 {
		return -1, errors.New("nothing to calculate")
	}
	if len(nodes) == 1 {
		if nodes[0].isOperatorType() {
			return -1, errors.New("single node input is operator")
		}
		return nodes[0].getValue(), nil
	}
	if len(nodes) == 2 {
		return -1, errors.New("we cannot reduce a two node input")
	}
	// Reduce beginning from right
	i := len(nodes)-3
	for i >= 0 {
		// We have a calculation to perform if we have an operator followed by two values
		if nodes[i].isOperatorType() && !nodes[i+1].isOperatorType() && !nodes[i+2].isOperatorType() {
			newValue := performCalculation(nodes[i].getOperator(), nodes[i+1].getValue(), nodes[i+2].getValue())
			newNodes := append(nodes[:i], CalculatorInt{value: newValue})
			newNodes = append(newNodes, nodes[i+3:]...)
			return reduceNodes(newNodes)
		}
		i--
	}
	return -1, errors.New("not a parsable input")
}

func performCalculation(operator string, num1 int, num2 int) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "/":
		//TODO: in a more production ready system we handle division by 0 and non-integer div properly
		if num2 == 0 {
			return 0
		}
		return num1 / num2
	case "*":
		return num1 * num2
	}
	return -1
}

func isPositiveInteger(s string) (int, bool) {
	val, boolean := isInteger(s)
	if boolean && val >= 0 {
		return val, true
	}
	return -1, false
}

func isInteger(s string) (int, bool) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return -1, false
	}
	return val, true
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}
