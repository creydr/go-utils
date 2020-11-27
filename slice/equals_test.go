package slice

import (
	"fmt"
	"testing"
)

func ExampleIsEqualStringSlice() {
	s1 := []string{"foo", "bar", "baz"}
	s2 := []string{"foo", "bar", "baz"}
	s3 := []string{"foobar", "foo"}
	fmt.Println(IsEqualStringSlice(s1, s2))
	fmt.Println(IsEqualStringSlice(s1, s3))
	fmt.Println(IsEqualStringSlice(s1, s1))
	fmt.Println(IsEqualStringSlice(nil, s2))
	fmt.Println(IsEqualStringSlice(nil, nil))
	// Output:
	// true
	// false
	// true
	// false
	// true
}

func ExampleIsEqualIntSlice() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 1}
	fmt.Println(IsEqualIntSlice(s1, s2))
	fmt.Println(IsEqualIntSlice(s1, s3))
	fmt.Println(IsEqualIntSlice(s1, s1))
	fmt.Println(IsEqualIntSlice(nil, s2))
	fmt.Println(IsEqualIntSlice(nil, nil))
	// Output:
	// true
	// false
	// true
	// false
	// true
}

func TestIsEqualStringSlice(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want bool
	}{
		{
			name: "Should return false on slices with different length",
			s1:   []string{"foo", "bar", "baz"},
			s2:   []string{"foo", "baz"},
			want: false,
		},
		{
			name: "Should return false on slices with same length but different values",
			s1:   []string{"foo", "bar", "baz"},
			s2:   []string{"foo", "baz", "foobar"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEqualStringSlice(tt.s1, tt.s2); got != tt.want {
				t.Errorf("IsEqualStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEqualIntSlice(t *testing.T) {
	tests := []struct {
		name string
		s1   []int
		s2   []int
		want bool
	}{
		{
			name: "Should return false on slices with different length",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2},
			want: false,
		},
		{
			name: "Should return false on slices with same length but different values",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 3, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEqualIntSlice(tt.s1, tt.s2); got != tt.want {
				t.Errorf("IsEqualIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
