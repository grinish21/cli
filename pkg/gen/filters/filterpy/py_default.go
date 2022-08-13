package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToDefaultString schema is nil")
	}
	if schema.Module == nil {
		return "", fmt.Errorf("ToDefaultString schema module is nil")
	}
	var text string
	if schema.IsArray {
		text = "[]"
	} else {
		switch schema.KindType {
		case model.TypeString:
			text = "\"\""
		case model.TypeInt:
			text = "0"
		case model.TypeFloat:
			text = "0.0"
		case model.TypeBool:
			text = "False"
		case model.TypeEnum:
			e := schema.Module.LookupEnum(schema.Type)
			if e == nil {
				return "", fmt.Errorf("ToDefaultString enum %s not found", schema.Type)
			}
			text = fmt.Sprintf("%s.%s", e.Name, e.Members[0].Name)
		case model.TypeStruct:
			s := schema.Module.LookupStruct(schema.Type)
			if s == nil {
				return "", fmt.Errorf("ToDefaultString struct %s not found", schema.Type)
			}
			text = "{}"
		case model.TypeInterface:
			i := schema.Module.LookupInterface(schema.Type)
			if i == nil {
				return "", fmt.Errorf("ToDefaultString interface %s not found", schema.Type)
			}
			text = "None"
		case model.TypeNull:
			text = "None"
		default:
			return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	}
	if text == "" {
		return "", fmt.Errorf("unknown type %s", schema.Type)
	}
	return text, nil
}

// cppDefault returns the default value for a type
func pyDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("called with nil node")
	}
	log.Debugf("pyDefault: %s", node.Name)
	return ToDefaultString(&node.Schema, prefix)
}