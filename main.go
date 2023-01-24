package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(v1 int, o string, v2 int) int {
	switch o {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	}
	panic("invalid operation:" + o)
}

func expand(exp []interface{}) []interface{} {
	for i := 0; i < len(exp)-2; i++ {
		o1, ok1 := exp[i].(string)
		v, ok2 := exp[i+1].(int)
		o2, ok3 := exp[i+2].(string)

		if ok1 && ok2 && ok3 && o1 == "(" && o2 == ")" {
			nexp := append(exp[:i+1])
			nexp[i] = v
			exp = append(nexp, exp[i+3:]...)
			i = i - 1
		}
	}
	return exp
}

func calc(exp []interface{}) []interface{} {
	for i := 0; i < len(exp)-2; i++ {
		v1, ok1 := exp[i].(int)
		o, ok2 := exp[i+1].(string)
		v2, ok3 := exp[i+2].(int)

		if ok1 && ok2 && ok3 {
			//fmt.Println(v1, o, v2)
			nexp := append(exp[:i+1])
			nexp[i] = calculate(v1, o, v2)
			exp = append(nexp, exp[i+3:]...)
			i = i - 1
			//fmt.Println(exp)
		}
	}
	return exp
}

func main() {
	tokens := []interface{}{}
	expression := "10 + ( 40 - ( 12 - 3 + 4 ) - ( 5 - 2 - 10 ) ) * 10 / 5"

	for _, s := range strings.Split(expression, " ") {
		if v, err := strconv.Atoi(s); err == nil {
			tokens = append(tokens, v)
		} else {
			tokens = append(tokens, s)
		}
	}
	tokens = calc(tokens)
	fmt.Println(tokens)
	tokens = expand(tokens)
	fmt.Println(tokens)
}
