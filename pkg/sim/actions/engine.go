package actions

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/spec"
)

// engine implements core.IEngine interface
var _ core.IEngine = (*Engine)(nil)

type Engine struct {
	eval *eval
	docs []*spec.ScenarioDoc
	core.Notifier
}

func NewEngine() *Engine {
	e := &Engine{
		eval: NewEval(),
		docs: make([]*spec.ScenarioDoc, 0),
	}
	e.init()
	return e
}

func (e *Engine) init() {
	e.eval.OnChange(func(symbol string, name string, value any) {
		e.EmitOnChange(symbol, name, value)
	})
	e.eval.OnSignal(func(symbol string, name string, args map[string]any) {
		e.EmitOnSignal(symbol, name, args)
	})
}

func (e *Engine) LoadScenario(source string, doc *spec.ScenarioDoc) error {
	doc.Source = source
	e.docs = append(e.docs, doc)
	for _, iface := range doc.Interfaces {
		if iface.Name == "" {
			return fmt.Errorf("interface %v has no name", iface)
		}
		log.Infof("registering interface %s\n", iface.Name)
	}
	return nil
}

func (a *Engine) UnloadScenario(source string) error {
	for i, d := range a.docs {
		if d.Source == source {
			a.docs = append(a.docs[:i], a.docs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("scenario %s not found", source)
}

func (e *Engine) HasInterface(ifaceId string) bool {
	for _, d := range e.docs {
		if d.GetInterface(ifaceId) != nil {
			return true
		}
	}
	return false
}

func (e *Engine) GetInterface(ifaceId string) *spec.InterfaceEntry {
	for _, d := range e.docs {
		if s := d.GetInterface(ifaceId); s != nil {
			return s
		}
	}
	return nil
}

// InvokeOperation invokes a operation of the interface.
func (e *Engine) InvokeOperation(symbol string, name string, args map[string]any) (any, error) {
	log.Infof("%s/%s invoke\n", symbol, name)
	iface := e.GetInterface(symbol)
	if iface == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	op := iface.GetOperation(name)
	if op == nil {
		return nil, fmt.Errorf("operation %s not found", name)
	}
	result, err := e.eval.EvalActions(symbol, op.Actions, iface.Properties)
	if err != nil {
		return nil, err
	}
	log.Infof("%s/%s result %v\n", symbol, name, result)
	return result, nil
}

// SetProperties sets the properties of the interface.
func (e *Engine) SetProperties(symbol string, props map[string]any) error {
	iface := e.GetInterface(symbol)
	if iface == nil {
		return fmt.Errorf("interface %s not found", symbol)
	}
	for name, value := range props {
		iface.Properties[name] = value
	}
	return nil
}

// FetchProperties returns a copy of the properties of the interface.
func (e *Engine) GetProperties(symbol string) (map[string]any, error) {
	iface := e.GetInterface(symbol)
	if iface == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	return iface.Properties, nil
}

func (e *Engine) HasSequence(sequencerId string) bool {
	for _, d := range e.docs {
		if d.GetSequence(sequencerId) != nil {
			return true
		}
	}
	return false
}

func (e *Engine) PlaySequence(sequencerId string) {
	for _, d := range e.docs {
		if s := d.GetSequence(sequencerId); s != nil {
			log.Printf("playing sequencer %s", sequencerId)
		}
	}
}
