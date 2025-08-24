package Lesson4Task1

import "VkEducation/FunctionExt"

type intArray FunctionExt.IntArray

func findMaxNumber(nums intArray) int {
	ret := 0
	leftIndex := 0
	rightIndex := len(nums) - 1
	for leftIndex <= rightIndex {
		if ret < nums[rightIndex] {
			ret = nums[rightIndex]
		}
		if ret < nums[leftIndex] {
			ret = nums[leftIndex]
		}
		rightIndex--
		leftIndex++
	}
	return ret
}
