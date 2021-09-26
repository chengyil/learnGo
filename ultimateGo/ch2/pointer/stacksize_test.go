package pointer

import "testing"

type bigAllocation [2048]int8
type smallAllocation [10]int8

func TestWithinStackSize(t *testing.T) {
	var allocation smallAllocation
	// allocating small allocation 1 byte by 10
	println(len(allocation), cap(allocation))

	var inStack int8

	// Value Semantic
	// And this is passing an array, which will copy the whole element in the array
	withinStackSize(&inStack, 10, allocation)
}

func withinStackSize(control *int8, depth int, allocate smallAllocation) {
	if depth < 0 {
		return
	}
	var i, j int8
	println(control, &i, &j, &allocate)
	withinStackSize(control, depth-1, allocate)
}

func TestExceedStackSize(t *testing.T) {
	var allocation bigAllocation
	// allocating small allocation 1 byte by 10
	println(len(allocation), cap(allocation))

	var inStack int8

	// Value Semantic
	// And this is passing an array, which will copy the whole element in the array
	// This will pass by Value, we will keep growing the stack
	exceedStackSize(&inStack, 10, allocation)
}

func exceedStackSize(control *int8, depth int, allocate bigAllocation) {
	if depth < 0 {
		return
	}
	var i, j int8
	println(control, &i, &j, &allocate)
	exceedStackSize(control, depth-1, allocate)
}
