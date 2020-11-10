package slice

import (
	"fmt"
)

func ExampleContainsString() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(ContainsString(s, "foo"))
	fmt.Println(ContainsString(s, "foobar"))
	fmt.Println(ContainsString(nil, "foo"))
	// Output:
	// true
	// false
	// false
}

func ExampleContainsInt() {
	s := []int{1, 2, 3}
	fmt.Println(ContainsInt(s, 1))
	fmt.Println(ContainsInt(s, 4))
	fmt.Println(ContainsInt(nil, 4))
	// Output:
	// true
	// false
	// false
}
