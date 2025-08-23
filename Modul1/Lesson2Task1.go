package Modul1

// CountNotEqual выполняет поиск элемента с помощью алгоритма двух указателей
func CountNotEqual(elements []int, nonEquivalent int) int {
	var rightIndex = len(elements) - 1
	var ret = 0
	for leftIndex := 0; leftIndex <= rightIndex; leftIndex++ {
		if elements[leftIndex] != nonEquivalent {
			ret++
		}
		if rightIndex != leftIndex && elements[rightIndex] != nonEquivalent {
			ret++
		}
		rightIndex--
	}
	return ret
}

// CountNotEqualLineal выполняет поиск элемента с помощью классического линейного поиска
func CountNotEqualLineal(elements []int, x int) int {
	count := 0
	for _, v := range elements {
		if v != x {
			count++
		}
	}
	return count
}
