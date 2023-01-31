package main

type CalculatorStack struct {
	stack  Stack
	proc   processor
	result int
}

func NewCalculatorStack() *CalculatorStack {
	return &CalculatorStack{stack: Stack{}, proc: processor{}}
}

func (c *CalculatorStack) Calculate(expression string) int {
	tokens := splitExpression(expression)
	for _, token := range tokens {
		c.proc.push(token)
		if c.proc.isReady() {
			c.calc()
		}
	}
	return c.proc.a.(int)
}

func (c *CalculatorStack) calc() {
	if c.proc.isExpression() {
		value := c.proc.calculate()
		if c.stack.length > 1 {
			c.proc.cba(value, c.stack.Pop(), c.stack.Pop())
			c.calc()
		} else if c.stack.length == 1 {
			c.proc.abc(c.stack.Pop(), value, nil)
		} else {
			c.proc.abc(value, nil, nil)
		}
	} else {
		c.stack.Push(c.proc.a)
		c.proc.abc(c.proc.b, c.proc.c, nil)
	}
}
