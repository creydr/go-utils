package slice

import (
	"fmt"
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
