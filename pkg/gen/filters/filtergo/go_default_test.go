package filtergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, method inputs, method outputs, signal inputs, struct fields
func TestDefaultFromIdl(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propFloat", "0.0"},
		{"test", "Test1", "propString", "\"\""},
		{"test", "Test1", "propBoolArray", "[]bool{}"},
		{"test", "Test1", "propIntArray", "[]int{}"},
		{"test", "Test1", "propFloatArray", "[]float64{}"},
		{"test", "Test1", "propStringArray", "[]string{}"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r := goDefault(prop, "")
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1Default"},
		{"test", "Test2", "propStruct", "Struct1{}"},
		{"test", "Test2", "propInterface", "nil"},
		{"test", "Test2", "propEnumArray", "[]Enum1{}"},
		{"test", "Test2", "propStructArray", "[]Struct1{}"},
		{"test", "Test2", "propInterfaceArray", "[]*Interface1{}"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r := goDefault(prop, "")
			assert.Equal(t, tt.rt, r)
		})
	}
}
