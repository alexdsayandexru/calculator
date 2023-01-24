package main

import "fmt"

type Calculator struct {
}

func (c Calculator) Calculate(expression string) {
	fmt.Println("Hello from calculator:", expression)
}
