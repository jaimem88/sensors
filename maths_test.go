package sensors

import (
	"reflect"
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"get the max value correctly",
			args{
				in: []int{0, 2, 4, 2, 4, 6, 3, 0, 6},
			},
			6,
		},
		{
			"zero on empty slice",
			args{
				in: []int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.in); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mode(t *testing.T) {
	type args struct {
		in []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			"get the correct mode",
			args{
				in: []float64{0.00, 2.33, 4.45, 2.1, 4.445, 4.45},
			},
			[]float64{4.45},
		},
		{
			"get the correct mode for more than 1 mode",
			args{
				in: []float64{0.00, 2.33, 4.45, 2.33, 4.445, 4.45},
			},
			[]float64{2.33, 4.45},
		},
		{
			"get the empty mode for non repeating values",
			args{
				in: []float64{0.00, 2.33, 4.5, 2.31, 4.445, 4.45},
			},
			[]float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mode(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_median(t *testing.T) {
	type args struct {
		in []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"get the correct median value (even)",
			args{
				in: []float64{0.00, 2.33, 4.45, 2.35},
			},
			2.35,
		},
		{
			"get the correct median value (odd)",
			args{
				in: []float64{0.00, 2.33, 4.45, 3.78, 7.4},
			},
			3.78,
		},
		{
			"get the zero median value",
			args{
				in: []float64{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := median(tt.args.in); got != tt.want {
				t.Errorf("median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_average(t *testing.T) {
	type args struct {
		in []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"get the average value",
			args{
				in: []float64{3.24, 2.43, 4.45, 2.35},
			},
			3.12,
		},
		{
			"get the zero value",
			args{
				in: []float64{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.in); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
