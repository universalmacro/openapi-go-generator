package gen

import "github.com/dave/jennifer/jen"

type StringEnum struct {
	Id     string
	Values []string
	Prefix string
}

func (se StringEnum) Statement() []*jen.Statement {
	var s []*jen.Statement
	s = append(s, jen.Type().Id(se.Id).String())
	for _, value := range se.Values {
		s = append(s, jen.Const().Id(se.Prefix+value).Id(se.Id).Op("=").Lit(value))
	}
	return s
}
