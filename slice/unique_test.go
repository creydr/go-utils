package slice

import (
	"fmt"
)

func ExampleUniqueStrings() {
	fmt.Printf("%#v\n", UniqueStrings([]string{"foo", "bar", "foo", "baz"}))
	fmt.Printf("%#v\n", UniqueStrings([]string{"foo", "bar", "baz"}))
	fmt.Printf("%#v\n", UniqueStrings(nil))
	// Output:
	// []string{"foo", "bar", "baz"}
	// []string{"foo", "bar", "baz"}
	// []string{}
}
