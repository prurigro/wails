package binding

import (
	"github.com/stretchr/testify/require"
	"github.com/wailsapp/wails/v2/internal/logger"
	"reflect"
	"strings"
	"testing"
)

type BindingTest struct {
	name        string
	structs     []interface{}
	exemptions  []interface{}
	want        string
	shouldError bool
}

func TestBindings_GenerateModels(t *testing.T) {

	tests := []BindingTest{
		SingleFieldTest,
		NestedFieldTest,
	}

	testLogger := &logger.Logger{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBindings(testLogger, tt.structs, tt.exemptions)
			for _, s := range tt.structs {
				err := b.Add(s)
				require.NoError(t, err)
			}
			got, err := b.GenerateModels()
			if (err != nil) != tt.shouldError {
				t.Errorf("GenerateModels() error = %v, shouldError %v", err, tt.shouldError)
				return
			}
			if !reflect.DeepEqual(strings.Fields(string(got)), strings.Fields(tt.want)) {
				t.Errorf("GenerateModels() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}
