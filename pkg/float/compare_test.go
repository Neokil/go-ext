package float_test

import (
	"testing"

	"github.com/Neokil/go-ext/pkg/float"
)

func Test_compareFloat(t *testing.T) {
	type args struct {
		f1        float64
		f2        float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 0.1 == 0.1 with precision 1",
			args: args{
				f1:        0.1,
				f2:        0.1,
				precision: 1,
			},
			want: true,
		},
		{
			name: "test 0.0001 == 0.0001 with precision 10",
			args: args{
				f1:        0.0001,
				f2:        0.0001,
				precision: 10,
			},
			want: true,
		},
		{
			name: "test 0.0001 == 0.0002 with precision 3",
			args: args{
				f1:        0.0001,
				f2:        0.0002,
				precision: 3,
			},
			want: true,
		},
		{
			name: "test 0.0001 == 0.0002 with precision 4",
			args: args{
				f1:        0.0001,
				f2:        0.0002,
				precision: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := float.Compare(tt.args.f1, tt.args.f2, tt.args.precision); got != tt.want {
				t.Errorf("compareFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
