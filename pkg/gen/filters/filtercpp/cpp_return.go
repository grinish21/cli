package filtercpp

import (
	"fmt"
	"reflect"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema) string {
	t := schema.Type
	text := ""
	switch t {
	case "string":
		text = "std::string"
	case "int":
		text = "int"
	case "float":
		text = "double"
	case "bool":
		text = "bool"
	default:
		if schema.Module == nil {
			log.Error().Msg("schema.Module is nil")
		}
		e := schema.Module.LookupEnum(t)
		if e != nil {
			text = e.Name
		}
		s := schema.Module.LookupStruct(t)
		if s != nil {
			text = s.Name
		}
		i := schema.Module.LookupInterface(t)
		if i != nil {
			text = fmt.Sprintf("%s*", i.Name)
		}
	}
	if schema.IsArray {
		text = fmt.Sprintf("std::vector<%s>", text)
	}
	return text
}

// cast value to TypedNode and deduct the cpp return type
func cppReturn(node reflect.Value) (reflect.Value, error) {
	p, ok := node.Interface().(model.ITypeProvider)
	if !ok {
		return reflect.ValueOf(""), fmt.Errorf("%s is not a schema provider", node.Type())
	}
	t := ToReturnString(p.GetSchema())
	return reflect.ValueOf(t), nil
}
