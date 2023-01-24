package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calc(exp []interface{}) []interface{} {
	for i := 0; i < len(exp)-2; i++ {
		v1, ok1 := exp[i].(int)
		v2, ok2 := exp[i+1].(string)
		v3, ok3 := exp[i+2].(int)

		if ok1 && ok2 && ok3 {
			fmt.Println(i, v1, v2, v3)
			nexp := append(exp[:i+1])
			nexp[i] = v1 - v3
			exp = append(nexp, exp[i+3:]...)
		}
	}
	return exp
}

func main() {
	tokens := []interface{}{}
	expression := "10 + ( 40 - ( 12 - 3 ) - ( 5 - 2 ) ) * 10 / 5"
	s1 := strings.Split(expression, " ")
	//fmt.Println(s1)
	//s2 := append(s1[:12], s1[15:]...)
	//fmt.Println(s2)
	for _, s := range s1 {
		if v, err := strconv.Atoi(s); err == nil {
			//fmt.Println(v)
			tokens = append(tokens, v)
		} else {
			tokens = append(tokens, s)
		}
	}
	fmt.Println(calc(tokens))
}
