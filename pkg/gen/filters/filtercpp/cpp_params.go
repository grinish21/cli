package filtercpp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func cppParams(node reflect.Value) (reflect.Value, error) {
	m, ok := node.Interface().(*model.Operation)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected method, got %s", node.Type())
	}
	var params []string
	for _, p := range m.Params {
		ps, err := ToParamString(p.GetSchema(), p.GetName())
		if err != nil {
			return reflect.Value{}, err
		}
		params = append(params, ps)
	}
	return reflect.ValueOf(strings.Join(params, ", ")), nil
}
