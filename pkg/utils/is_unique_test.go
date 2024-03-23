package utils

import "testing"

func TestIsUnique(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		args []T
		want bool
	}
	tests := []testCase[string]{
		{
			name: "for strings 1 - not unique",
			args: []string{"a", "a", "b", "c"},
			want: false,
		},
		{
			name: "for strings 2 - unique",
			args: []string{"q", "a", "b", "c"},
			want: true,
		},
		{
			name: "for strings 3 - with empty str",
			args: []string{"", "a", "b", "c"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnique(tt.args); got != tt.want {
				t.Errorf("IsUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
