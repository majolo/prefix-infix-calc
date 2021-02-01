package main

// Node is an interface to tokenize operators and integers.
type Node interface {
	isOperatorType() bool
	getValue() int
	getOperator() string
}

type CalculatorInt struct {
	value int
}

func (c CalculatorInt) isOperatorType() bool {
	return false
}

func (c CalculatorInt) getValue() int {
	return c.value
}

func (c CalculatorInt) getOperator() string {
	return ""
}

type CalculatorOperator struct {
	operator string
}

func (c CalculatorOperator) isOperatorType() bool {
	return true
}

func (c CalculatorOperator) getValue() int {
	return -1
}

func (c CalculatorOperator) getOperator() string {
	return c.operator
}

