package Lesson3Task2

import "VkEducation/FunctionExt"

type intArray FunctionExt.IntArray

func DeleteExtraNums(targetArray intArray, val2Delete int) intArray {
	currentIndex := 0
	for _, v := range targetArray {
		if v != val2Delete {
			targetArray[currentIndex] = v
			currentIndex++
		}
	}
	return targetArray[:currentIndex]
}
