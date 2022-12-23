package main

import "testing"

func TestTotal(t *testing.T) {
	type data struct {
		values []int
	}

	tests := []struct {
		name string
		vals data
		want int
	}{
		{
			name: "first test case Total",
			vals: data{
				[]int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "second test case Total",
			vals: data{
				[]int{2, 3, 4},
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Total(tt.vals.values...); got != tt.want {
				t.Errorf("Total() = %v, want %v", got, tt.want)
			}
		})
	}
}
