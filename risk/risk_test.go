package risk

import (
	"testing"
)

func Test_computeAverage(t *testing.T) {
	type args struct {
		values []*Value
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Low average",
			args{[]*Value{&Value{1, ""}, &Value{2, ""}, &Value{2, ""}, &Value{2, ""}}},
			1.75,
		},
		{
			"Medium average",
			args{[]*Value{&Value{5, ""}, &Value{7, ""}, &Value{6, ""}, &Value{5, ""}}},
			5.75,
		},
		{
			"High average",
			args{[]*Value{&Value{7, ""}, &Value{9, ""}, &Value{9, ""}, &Value{8, ""}}},
			8.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeAverage(tt.args.values); got != tt.want {
				t.Errorf("computeAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_displayAverage(t *testing.T) {
	type args struct {
		values float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Low rating",
			args{1.75},
			"{set:cellbgcolor:lightgreen}1.75 (LOW)",
		},
		{
			"Medium rating",
			args{5.75},
			"{set:cellbgcolor:yellow}5.75 (MEDIUM)",
		},
		{
			"High rating",
			args{8.25},
			"{set:cellbgcolor:red}8.25 (HIGH)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := displayAverage(tt.args.values); got != tt.want {
				t.Errorf("displayAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
