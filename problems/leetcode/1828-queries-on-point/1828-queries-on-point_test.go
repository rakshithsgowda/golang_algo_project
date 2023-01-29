package leetcode

import (
	"reflect"
	"testing"
)

func Test_countPoints(t *testing.T) {
	// t.Parallel()
	type args struct {
		points  [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{
				points:  [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
				queries: [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			},
			want: []int{3, 2, 2},
		},
		{
			name: "example 2",
			args: args{
				points:  [][]int{{1, 1}, {2, 2}, {4, 4}, {5, 5}},
				queries: [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}},
			},
			want: []int{2, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := countPoints(tt.args.points, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
