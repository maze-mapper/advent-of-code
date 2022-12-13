package day13

import (
	"testing"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		left, right []any
		want        int
	}{
		{left: []any{1, 1, 3, 1, 1}, right: []any{1, 1, 5, 1, 1}, want: correctOrder},
		{left: []any{[]any{1}, []any{2, 3, 4}}, right: []any{[]any{1}, 4}, want: correctOrder},
		{left: []any{9}, right: []any{[]any{8, 7, 6}}, want: wrongOrder},
		{left: []any{[]any{4, 4}, 4, 4}, right: []any{[]any{4, 4}, 4, 4, 4}, want: correctOrder},
		{left: []any{7, 7, 7, 7}, right: []any{7, 7, 7}, want: wrongOrder},
		{left: []any{[]any{[]any{}}}, right: []any{[]any{}}, want: wrongOrder},
		{left: []any{1, []any{2, []any{3, []any{4, []any{5, 6, 7}}}}, 8, 9}, right: []any{1, []any{2, []any{3, []any{4, []any{5, 6, 0}}}}, 8, 9}, want: wrongOrder},
	}

	for _, tc := range tests {
		t.Run("TODO", func(t *testing.T) {
			got := compare(tc.left, tc.right)
			if got != tc.want {
				t.Errorf("compare(%#v, %#v) = %d, want %d", tc.left, tc.right, got, tc.want)
			}
		})
	}
}
