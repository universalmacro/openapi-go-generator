package gen

import (
	"github.com/dave/jennifer/jen"
)

type GStruct struct {
	Id     string
	Fields []Field
}

type Field struct {
	Variable
	Tags []string
}

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

func (ps Outputs) Apply(c *jen.Statement) *jen.Statement {
	if len(ps) == 0 {
		return c
	}
	if len(ps) == 1 {
		if ps[0].Type.IsPointer {
			c.Op("*")
		}
		return c.Id(ps[0].Type.Name)
	}
	c.Params(
		c.String(), c.String())
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
	Name string
	Type Type
	Init *string
}

func (v Variable) Parameter() jen.Code {
	c := jen.Id(v.Name)
	if v.Type.IsPointer {
		c = c.Op("*")
	}
	c.Id(v.Type.Name)
	return c
}

type Type struct {
	Name      string
	IsPointer bool
}
