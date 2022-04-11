package filtercpp

import (
	"fmt"
	"objectapi/pkg/model"
	"reflect"
	"strings"
)

func ToParamString(t string) string {
	isArray := strings.HasSuffix(t, "[]")
	if isArray {
		t = t[:len(t)-2]
	}
	s := ""
	switch t {
	case "string":
		s = "std::string"
	case "int":
		s = "int"
	case "float":
		s = "double"
	case "bool":
		s = "bool"
	default:
		s = t
	}
	if isArray {
		s = fmt.Sprintf("std::vector<%s>", s)
	}
	return s
}

func cppParam(node reflect.Value) (reflect.Value, error) {
	schema := node.Interface().(*model.TypedNode).Schema
	t := ToParamString(schema.Type)
	return reflect.ValueOf(t), nil
}
