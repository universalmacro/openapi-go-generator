package gen

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
)

func TestFunc(t *testing.T) {
	fun := Func{Id: "Test"}
	fun.AddParameters(Variable{Id: "a", Type: Type{Id: "string"}}, Variable{Id: "b", Type: Type{Id: "string"}})
	fun.AddOutputs(Variable{Type: Type{Id: "int", IsPointer: true}}, Variable{Type: Type{Id: "int"}})
	f := jen.NewFile("models")
	st := Struct{Id: "HelloWorld"}
	st.AddFields(Field{Tags: map[string]string{"json": "good"}, Variable: Variable{Id: "A", Type: Type{Id: "string"}}}, Field{Variable: Variable{Id: "B", Type: Type{Id: "string"}}})
	method := Method{Func: Func{
		Id: "HelloWorld",
	}, SelfId: "h"}
	method.AddStatements(jen.Var().Id("good").String())
	st.AddMethods(method)
	f.Add(fun.Statement())
	for _, s := range st.Statement() {
		f.Add(s)
	}
	fmt.Printf("%#v", f)
}
