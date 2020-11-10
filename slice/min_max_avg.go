package slice

// MinInt returns the smallest integer value of the given slice
func MinInt(slice []int) int {
	if len(slice) == 0 {
		panic("Can't pass an empty or nil slice to MinInt function")
	}

	return MinIntVar(slice[0], slice[1:]...)
}

// MinIntVar returns the smallest integer value of the values
func MinIntVar(val int, vals ...int) int {
	min := val
	for _, i := range vals {
		if min > i {
			min = i
		}
	}

	return min
}

// MaxInt returns the biggest integer value of the given slice
func MaxInt(slice []int) int {
	if len(slice) == 0 {
		panic("Can't pass an empty or nil slice to MaxInt function")
	}

	return MaxIntVar(slice[0], slice[1:]...)
}

// MaxIntVar returns the biggest integer value of the values
func MaxIntVar(val int, vals ...int) int {
	min := val
	for _, i := range vals {
		if min < i {
			min = i
		}
	}

	return min
}

// AvgInt returns the average value of the given slice
func AvgInt(slice []int) float32 {
	if len(slice) == 0 {
		panic("Can't pass an empty or nil slice to AvgInt function")
	}

	return AvgIntVar(slice[0], slice[1:]...)
}

// AvgIntVar returns the average value of the values
func AvgIntVar(val int, vals ...int) float32 {
	sum := val
	for _, i := range vals {
		sum += i
	}

	return float32(sum) / float32(len(vals)+1)
}
