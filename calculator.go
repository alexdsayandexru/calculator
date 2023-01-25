package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
}

func (c Calculator) Calculate(expression string) (int, bool) {
	tokens := splitExpression(expression)
	tokens = calculateTokens(tokens)
	value, ok := tokens[0].(int)
	if ok {
		return value, true
	}
	return 0, false
}

func printTokens(tokens []interface{}){
	for i, token := range tokens {
		value, ok := token.(int)
		operator, _ := token.(string)

		if(ok) {
			fmt.Println(i, "v:", value)
		} else {
			fmt.Println(i, "o:", operator)
		}
	}
}

func calculate(left int, right int, operator string) int {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	panic("invalid operator:" + operator)
}

func splitExpression (expression string) []interface{} {
	tokens := []interface{}{}
	for _, s := range strings.Split(expression, " ") {
		if v, err := strconv.Atoi(s); err == nil {
			tokens = append(tokens, v)
		} else {
			tokens = append(tokens, s)
		}
	}
	return tokens
}

func isSimpleExpression(tokens []interface{}, i int) (bool, int, int, string) {
	left, ok_left := tokens[i].(int)
	operator, ok_operator := tokens[i+1].(string)
	right, ok_right := tokens[i+2].(int)

	if ok_left && ok_operator && ok_right {
		return true, left, right, operator
	}
	return false, left, right, operator
}

func isBrackets(tokens []interface{}, i int) (bool, int) {
	if len(tokens) == 3 {
		i = 0
	}
	o1, ok1 := tokens[i].(string)
	value, ok2 := tokens[i+1].(int)
	o2, ok3 := tokens[i+2].(string)

	if ok1 && ok2 && ok3 && o1 == "(" && o2 == ")" {
		return true, value
	}
	return false, value
}

func calculateTokens(tokens []interface{}) []interface{} {
	for i := 0; i < len(tokens)-2; i++ {
		if len(tokens) >=3 {
			ok, left, right, operator := isSimpleExpression(tokens, i)
			if ok {
				value := calculate(left, right, operator)
				tokens = transformTokens(tokens, i, value)
				i = i - 1
			}
		}
		if len(tokens) >=3 {
			ok, value := isBrackets(tokens, i)
			if ok {
				tokens = transformTokens(tokens, i, value)
				i = 0
			}
		}
		if len(tokens) == 3 {
			tokens = calculateTokens(tokens)
		}
	}
	return tokens
}

func transformTokens(tokens []interface{}, i int, value int)  []interface{} {
	new_tokens := append(tokens[:i+1])
	new_tokens[i] = value
	tokens = append(new_tokens, tokens[i+3:]...)
	return tokens
}
