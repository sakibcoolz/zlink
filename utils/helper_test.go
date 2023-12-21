package utils

import (
	"reflect"
	"strings"
	"testing"
)

func TestIntToStringEncode(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Pass1",
			args: args{
				ints: []int{1, 2, 3, 4, 5, 6},
			},
			want: []string{"L", "A", "A"},
		},
		{
			name: "Pass2",
			args: args{
				ints: []int{1, 4, 3, 5, 5, 6, 9},
			},
			want: []string{"N", "B", "A", "Q"},
		},
		{
			name: "Pass3",
			args: args{
				ints: []int{1, 1, 8, 2, 4, 9, 9, 1},
			},
			want: []string{"K", "F", "E", "H"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToStringEncode(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToStringEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToEncode(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Pass1",
			args: args{
				val: 123456,
			},
			want: "LAA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToEncode(tt.args.val); !strings.HasPrefix(got, tt.want) {
				t.Errorf("ToEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitToGigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Pass",
			args: args{n: 123},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitToGigits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitToGigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
