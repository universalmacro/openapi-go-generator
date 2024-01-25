package openapi

import "strings"

func generateTypeString(s Schema) string {
	if s.Type != nil {
		switch *s.Type {
		case "string":
			return "string"
		case "integer":
			if s.Format != nil && *s.Format == "int64" {
				return "int64"
			}
			return "int32"
		case "boolean":
			return "bool"
		case "number":
			if s.Format != nil && *s.Format == "double" {
				return "float64"
			}
			return "float32"
		case "array":
			if s.Items == nil {
				panic("array without items")
			}
			return "[]" + generateTypeString(*s.Items)
		}
		return *s.Type
	}
	if s.Ref != nil {
		return refToType(*s.Ref)
	}
	panic("type string")
}

func refToType(ref string) string {
	s := strings.Split(ref, "/")
	return s[len(s)-1]
}

func capitalize(str string) string {
	bs := []byte(str)
	if len(bs) == 0 {
		return ""
	}
	bs[0] = byte(bs[0] - 32)
	return string(bs)
}

func isRequired(name string, required *[]string) bool {
	if required == nil {
		return false
	}
	for _, r := range *required {
		if r == name {
			return true
		}
	}
	return false
}

func pathConvert(path string) string {
	return strings.ReplaceAll(strings.ReplaceAll(path, "{", ":"), "}", "")
}
