package main

import "fmt"

func main() {
	fmt.Println("")

	//expression := "10 + ( 40 - ( 12 - ( 3 + 4 ) ) - ( 5 - 2 - 10 ) ) * ( 5 + ( 10 / 5 ) )"
	//expression := "( 10 + 40 )"
	//expression := "2 + ( 100 + 5 ) * 3"

	/*calc1 := NewCalculatorStack()
	fmt.Println(calc1.Calculate("10 + 4"))

	calc2 := NewCalculatorStack()
	fmt.Println(calc2.Calculate("( 34 - 10 )"))

	calc3 := NewCalculatorStack()
	fmt.Println(calc3.Calculate("2 * ( 10 + 7 )"))

	calc4 := NewCalculatorStack()
	fmt.Println(calc4.Calculate("2 + ( 18 + 2 ) * 2"))*/

	calc5 := NewCalculatorStack()
	fmt.Println(calc5.Calculate("10 + ( 40 - ( 12 - ( 3 + 4 ) ) - ( 5 - 2 - 10 ) ) * ( 5 + ( 10 / 5 ) )"))
}
