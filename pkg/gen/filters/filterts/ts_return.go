package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeString:
		text = "string"
	case model.TypeInt:
		text = "number"
	case model.TypeFloat:
		text = "number"
	case model.TypeBool:
		text = "boolean"
	case model.TypeEnum:
		e := schema.Module.LookupEnum(schema.Type)
		if e == nil {
			return "", fmt.Errorf("ToReturnString enum %s not found", schema.Type)
		}
		text = fmt.Sprintf("%s%s", prefix, e.Name)
	case model.TypeStruct:
		s := schema.Module.LookupStruct(schema.Type)
		if s == nil {
			return "", fmt.Errorf("ToReturnString struct %s not found", schema.Type)
		}
		text = fmt.Sprintf("%s%s", prefix, s.Name)
	case model.TypeInterface:
		i := schema.Module.LookupInterface(schema.Type)
		if i == nil {
			return "", fmt.Errorf("ToReturnString interface %s not found", schema.Type)
		}
		text = fmt.Sprintf("%s%s", prefix, i.Name)
	case model.TypeNull:
		text = "void"
	default:
		return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		text = fmt.Sprintf("%s[]", text)
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func tsReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("tsReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}