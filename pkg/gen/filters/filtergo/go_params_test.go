package filtergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "param1 bool"},
		{"test", "Test3", "opInt", "param1 int64"},
		{"test", "Test3", "opFloat", "param1 float64"},
		{"test", "Test3", "opString", "param1 string"},
		{"test", "Test3", "opBoolArray", "param1 []bool"},
		{"test", "Test3", "opIntArray", "param1 []int64"},
		{"test", "Test3", "opFloatArray", "param1 []float64"},
		{"test", "Test3", "opStringArray", "param1 []string"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, meth)
			r, err := goParams("", meth.Params)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestParamsSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "param1 Enum1"},
		{"test", "Test4", "opStruct", "param1 Struct1"},
		{"test", "Test4", "opInterface", "param1 *Interface1"},
		{"test", "Test4", "opEnumArray", "param1 []Enum1"},
		{"test", "Test4", "opStructArray", "param1 []Struct1"},
		{"test", "Test4", "opInterfaceArray", "param1 []*Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := goParams("", prop.Params)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestParamsMultiple(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "param1 bool, param2 bool"},
		{"test", "Test5", "opIntInt", "param1 int64, param2 int64"},
		{"test", "Test5", "opFloatFloat", "param1 float64, param2 float64"},
		{"test", "Test5", "opStringString", "param1 string, param2 string"},
		{"test", "Test5", "opEnumEnum", "param1 Enum1, param2 Enum1"},
		{"test", "Test5", "opStructStruct", "param1 Struct1, param2 Struct1"},
		{"test", "Test5", "opInterfaceInterface", "param1 *Interface1, param2 *Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := goParams("", prop.Params)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestParamsWithErrors(t *testing.T) {
	s, err := goParams("", nil)
	assert.Error(t, err)
	assert.Equal(t, "", s)
}
