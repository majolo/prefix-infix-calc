package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type CalculationTest struct {
	expectedErr bool
	calc string
	expectedResult int
}

func TestPrefixCalculator(t *testing.T) {
	tests := []CalculationTest{
		{
			expectedErr:    false,
			calc:           "3",
			expectedResult: 3,
		},
		{
			expectedErr:    false,
			calc:           "+ 1 2",
			expectedResult: 3,
		},
		{
			expectedErr:    false,
			calc:           "+ 1 * 2 3",
			expectedResult: 7,
		},
		{
			expectedErr:    false,
			calc:           "+ * 1 2 3",
			expectedResult: 5,
		},
		{
			expectedErr:    false,
			calc:           "- / 10 + 1 1 * 1 2",
			expectedResult: 3,
		},
		{
			expectedErr:    false,
			calc:           "- 0 3",
			expectedResult: -3,
		},
		{
			expectedErr:    false,
			calc:           "/ 3 2",
			expectedResult: 1,
		},
		{
			expectedErr:    true,
			calc:           "/ 3 2 2",
			expectedResult: -1,
		},
	}
	for _, test := range tests {
		val, err := prefixCalculator(test.calc)
		if test.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, test.expectedResult, val, fmt.Sprintf("failed test with input %s", test.calc))
	}
}

func TestInfixCalculator(t *testing.T) {
	tests := []CalculationTest{
		{
			expectedErr:    false,
			calc:           "( 1 + 2 )",
			expectedResult: 3,
		},
		{
			expectedErr:    false,
			calc:           "( 1 + ( 2 * 3 ) )",
			expectedResult: 7,
		},
		{
			expectedErr:    false,
			calc:           "( ( 1 * 2 ) + 3 )",
			expectedResult: 5,
		},
		{
			expectedErr:    false,
			calc:           "( ( ( 1 + 1 ) / 10 ) - ( 1 * 2 ) )",
			expectedResult: -2,
		},
		{
			expectedErr:    true,
			calc:           "/ 3 2 2",
			expectedResult: -1,
		},
	}
	for _, test := range tests {
		val, err := infixCalculator(test.calc)
		if test.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, test.expectedResult, val, fmt.Sprintf("failed test with input %s", test.calc))
	}
}

func TestIsPositiveInteger(t *testing.T) {
	val, boolean := isPositiveInteger("10")
	require.True(t, boolean)
	require.Equal(t, val, 10)
	_, boolean = isPositiveInteger("kheiron")
	require.False(t, boolean)
	_, boolean = isPositiveInteger("-1")
	require.False(t, boolean)
	_, boolean = isPositiveInteger("10.1")
	require.False(t, boolean)
}

func TestIsOperator(t *testing.T) {
	require.True(t, isOperator("*"))
	require.False(t, isOperator("//"))
}
