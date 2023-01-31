package main

type processor struct {
	a interface{}
	b interface{}
	c interface{}
}

func (p *processor) abc(aa interface{}, bb interface{}, cc interface{}) {
	p.a = aa
	p.b = bb
	p.c = cc
}

func (p *processor) cba(cc interface{}, bb interface{}, aa interface{}) {
	p.a = aa
	p.b = bb
	p.c = cc
}

func (p *processor) isReady() bool {
	if p.a != nil && p.b != nil && p.c != nil {
		return true
	}
	return false
}

func (p *processor) push(token interface{}) {
	if p.a == nil {
		p.a = token
	} else if p.b == nil {
		p.b = token
	} else if p.c == nil {
		p.c = token
	} else {
		panic("There is no free register")
	}
}

func (p *processor) calculate() int {
	switch p.b {
	case "+":
		return p.a.(int) + p.c.(int)
	case "-":
		return p.a.(int) - p.c.(int)
	case "*":
		return p.a.(int) * p.c.(int)
	case "/":
		return p.a.(int) / p.c.(int)
	}

	return p.b.(int)
}

func (p *processor) isExpression() bool {
	_, ok_a := p.a.(int)
	_, ok_b := p.b.(string)
	_, ok_c := p.c.(int)

	if ok_a && ok_b && ok_c {
		return true
	}

	a, ok_a := p.a.(string)
	_, ok_b = p.b.(int)
	c, ok_c := p.c.(string)

	if ok_a && ok_b && ok_c && a == "(" && c == ")" {
		return true
	}

	return false
}
