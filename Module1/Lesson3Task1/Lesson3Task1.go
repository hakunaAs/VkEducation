package Lesson3Task1

import (
	"VkEducation/FunctionExt"
)

type intArray FunctionExt.IntArray

// SortZero
// Сортирует нули в конец, а порядок чисел неизменен
func SortZero(grades intArray) intArray {
	ret := make(intArray, len(grades))
	currentRetIndex := 0
	zeroCount := 0

	for _, value := range grades {
		if value == 0 {
			zeroCount++
		} else {
			ret[currentRetIndex] = value
			currentRetIndex++
		}
	}

	for i := 0; i < zeroCount; i++ {
		ret[currentRetIndex] = 0
		currentRetIndex++
	}

	return ret
}
