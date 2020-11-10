package slice

import (
	"fmt"
)

func ExampleMinInt() {
	fmt.Println(MinInt([]int{1, 2, 3}))
	fmt.Println(MinInt([]int{1, -1, 3}))
	fmt.Println(MinInt([]int{2}))
	// Output:
	// 1
	// -1
	// 2
}

func ExampleMinIntVar() {
	fmt.Println(MinIntVar(1, 2, 3))
	fmt.Println(MinIntVar(1, -1, 3))
	fmt.Println(MinIntVar(2))
	// Output:
	// 1
	// -1
	// 2
}

func ExampleMaxInt() {
	fmt.Println(MaxInt([]int{1, -1, 3}))
	fmt.Println(MaxInt([]int{1, 2, 3}))
	fmt.Println(MaxInt([]int{2}))
	// Output:
	// 3
	// 3
	// 2
}

func ExampleMaxIntVar() {
	fmt.Println(MaxIntVar(1, 2, 3))
	fmt.Println(MaxIntVar(1, -1, 3))
	fmt.Println(MaxIntVar(2))
	// Output:
	// 3
	// 3
	// 2
}

func ExampleAvgInt() {
	fmt.Println(AvgInt([]int{1, -1, 3}))
	fmt.Println(AvgInt([]int{1, 2, 3}))
	fmt.Println(AvgInt([]int{2}))
	// Output:
	// 1
	// 2
	// 2
}

func ExampleAvgIntVar() {
	fmt.Println(AvgIntVar(1, -1, 3))
	fmt.Println(AvgIntVar(1, 2, 3))
	fmt.Println(AvgIntVar(2))
	// Output:
	// 1
	// 2
	// 2
}
