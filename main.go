package main

import (
	"VkEducation/Module1"
	"fmt"
	"time"
)

func testByTime(countFunc func(elements []int, nonEquivalent int) int) {
	data := make([]int, 1<<30) // ~1 млд элементов
	for i := range data {
		data[i] = i & 7
	}

	start := time.Now()
	result := countFunc(data, 3)
	elapsed := time.Since(start)

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Elapsed: %s\n", elapsed)
}

func main() {
	//Выполнено за 446.2497ms
	testByTime(Module1.CountNotEqual)
	//Выполнено за 602.3725ms
	testByTime(Module1.CountNotEqualLineal)
}
