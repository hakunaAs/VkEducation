package FunctionExt

type IntArray []int

func swap(toSwap *IntArray, startIndex int, endIndex int) {
	swapItem := (*toSwap)[endIndex]
	(*toSwap)[endIndex] = (*toSwap)[startIndex]
	(*toSwap)[startIndex] = swapItem
}
