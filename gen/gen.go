package gen

import (
	"github.com/dave/jennifer/jen"
)

type GFunc struct {
	Id         string
	Parameters Parameters
	Outputs    Outputs
}

func (gf *GFunc) AddOutputs(p ...Variable) *GFunc {
	gf.Outputs = append(gf.Outputs, p...)
	return gf
}

func (gf *GFunc) AddParameters(p ...Variable) *GFunc {
	gf.Parameters = append(gf.Parameters, p...)
	return gf
}

func (f GFunc) Statement() *jen.Statement {
	s := jen.Func().Id(f.Id).Params(f.Parameters.Code()...)
	f.Outputs.Apply(s)
	s = s.Block()
	return s
}

type Outputs []Variable

func (outputs Outputs) Apply(c *jen.Statement) *jen.Statement {
	if len(outputs) == 0 {
		return c
	}
	if len(outputs) == 1 {
		if outputs[0].Type.IsPointer {
			c.Op("*")
		}
		return c.Id(outputs[0].Type.Id)
	}
	c.Op("(")
	for _, o := range outputs {
		if o.Type.IsPointer {
			c.Op("*")
		}
		c.Id(o.Type.Id)
		c.Op(",")
	}
	c.Op(")")
	return c
}

type Parameters []Variable

func (ps Parameters) Code() []jen.Code {
	var p []jen.Code
	for _, param := range ps {
		p = append(p, param.Parameter())
	}
	return p
}

type Variable struct {
	Id   string
	Type Type
	Init *string
}

func (v Variable) Parameter() jen.Code {
	c := jen.Id(v.Id)
	if v.Type.IsPointer {
		c = c.Op("*")
	}
	c.Id(v.Type.Id)
	return c
}

type Type struct {
	Id        string
	IsPointer bool
}
