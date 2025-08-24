package Lesson2Task1

import (
	"VkEducation/FunctionExt"
	"testing"
	"time"
)

func TestAlgorithm(t *testing.T) {
	cases := []struct {
		name    string
		arr     []int
		toCount int
		want    int
	}{
		{
			"fromExample",
			FunctionExt.IntArray{7, 8, 1, 3, 11, 3, 9, 4, 3},
			3,
			6,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := CountNotEqual(tc.arr, tc.toCount)
			if res != tc.want {
				t.Fatal("Неверно подсчитано количество не равных значений в массиве алгоритмом двух указателей!",
					"\nМассив: ", tc.arr, "\nПроверяемое число: ", tc.toCount,
					"\nФактический результат: ", res, "\nОжидаемый результат: ", tc.want)
			}
			res = CountNotEqualLineal(tc.arr, tc.toCount)
			if res != tc.want {
				t.Fatal("\nНеверно подсчитано количество не равных значений в массиве линейным алгоритмом!",
					"\nМассив: ", tc.arr, "\nПроверяемое число: ", tc.toCount,
					"\nФактический результат: ", res, "\nОжидаемый результат: ", tc.want)
			}
		})
	}
}

func startBenchmark(b *testing.B, name string, data intArray, countFunc func(elements intArray, nonEquivalent int) int) {
	b.Run(name, func(b *testing.B) {
		b.ResetTimer()
		start := time.Now()
		for i := 0; i < b.N; i++ {
			_ = countFunc(data, 3)
		}
		elapsed := time.Since(start)
		b.ReportMetric(float64(elapsed.Milliseconds())/float64(b.N), "ms/op")
	})
}

func BenchmarkAlgorithm(b *testing.B) {
	data := make(intArray, 1<<30) // ~1 млд элементов (нужно минимум 16 Гб оперативной памяти)
	for i := range data {
		data[i] = i & 7
	}
	startBenchmark(b, "twoPointers", data, CountNotEqual)
	startBenchmark(b, "linear", data, CountNotEqualLineal)
}
