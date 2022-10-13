package filterqt

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
		{"test", "Test3", "opBool", "bool param1"},
		{"test", "Test3", "opInt", "int param1"},
		{"test", "Test3", "opFloat", "double param1"},
		{"test", "Test3", "opString", "const QString &param1"},
		{"test", "Test3", "opBoolArray", "const QList<bool> &param1"},
		{"test", "Test3", "opIntArray", "const QList<int> &param1"},
		{"test", "Test3", "opFloatArray", "const QList<double> &param1"},
		{"test", "Test3", "opStringArray", "const QList<QString> &param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := qtParams("", meth.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "Enum1 param1"},
		{"test", "Test4", "opStruct", "const Struct1 &param1"},
		{"test", "Test4", "opInterface", "Interface1 *param1"},
		{"test", "Test4", "opEnumArray", "const QList<Enum1> &param1"},
		{"test", "Test4", "opStructArray", "const QList<Struct1> &param1"},
		{"test", "Test4", "opInterfaceArray", "const QList<Interface1*> &param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := qtParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsMultiple(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "bool param1, bool param2"},
		{"test", "Test5", "opIntInt", "int param1, int param2"},
		{"test", "Test5", "opFloatFloat", "double param1, double param2"},
		{"test", "Test5", "opStringString", "const QString &param1, const QString &param2"},
		{"test", "Test5", "opEnumEnum", "Enum1 param1, Enum1 param2"},
		{"test", "Test5", "opStructStruct", "const Struct1 &param1, const Struct1 &param2"},
		{"test", "Test5", "opInterfaceInterface", "Interface1 *param1, Interface1 *param2"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := qtParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
