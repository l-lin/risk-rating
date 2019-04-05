package risk

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want *Value
	}{
		{
			"Nominal case",
			args{"5 (loss of goodwill)"},
			&Value{5, "loss of goodwill"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
