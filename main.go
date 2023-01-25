package main

import "fmt"


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


func main() {
	fmt.Println("")
	
	expression := "10 + ( 40 - ( 12 - ( 3 + 4 ) ) - ( 5 - 2 - 10 ) ) * ( 5 + ( 10 / 5 ) )"
	//expression := "( 10 + 40 )"
	//expression := "100"

	calc := Calculator{}
	fmt.Println(calc.Calculate(expression))
}
