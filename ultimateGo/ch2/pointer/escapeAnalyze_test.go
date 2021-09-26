package pointer

import "testing"

const Data int = 1

func TestStack(t *testing.T) {
	allocateInStack(2)
}

func allocateInStack(depth int) int8 {
	if depth < 0 {
		var ret int8
		// ret is allocated in stack
		// Notice the value of the memory allocated
		// It is the same as i, j
		println("Return By Value Semantic", &ret)
		return ret

	}
	// i and j are allocated in stack
	var i, j, k uint8 = 0, 1<<8 - 1, 0
	ret := allocateInStack(depth - 1)
	// The new value is created in the same stack as i, j
	print("Depth: ", depth, ", ")
	println("Receive Result with Value Semantic", &i, &j, &k, &ret, j)
	return ret
}

func TestHeap(t *testing.T) {
	allocateInHeap(2)
}

func allocateInHeap(depth int) *int8 {
	if depth < 0 {
		var ret, stack int8

		// ret is allocated in heap
		// Notice the value of the memory allocated
		// It is different from stack
		println("Return By Pointer Semantic", &stack, &ret)
		return &ret

	}
	// i j k are allocated in stack
	var i, j, k uint8 = 0, 1<<8 - 1, 0
	ret := allocateInHeap(depth - 1)
	// The new value is created differently to i, j
	print("Depth: ", depth, ", ")
	println("Receive Result with Pointer Semantic", &i, &j, &k, ret, j)
	return ret
}
