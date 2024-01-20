package gen

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
)

func TestFunc(t *testing.T) {
	fun := GFunc{Id: "Test"}
	fun.AddParameters(Variable{Id: "a", Type: Type{Id: "string"}}, Variable{Id: "b", Type: Type{Id: "string"}})
	fun.AddOutputs(Variable{Type: Type{Id: "int", IsPointer: true}}, Variable{Type: Type{Id: "int"}})
	f := jen.NewFile("models")
	st := GStruct{Id: "HelloWorld"}
	st.AddFields(Field{Tags: map[string]string{"json": "good"}, Variable: Variable{Id: "A", Type: Type{Id: "string"}}}, Field{Variable: Variable{Id: "B", Type: Type{Id: "string"}}})
	st.AddMethods(Method{Method: fun, SelfId: "h"})
	f.Add(fun.Statement())
	f.Add(st.Statement())
	fmt.Printf("%#v", f)
}
