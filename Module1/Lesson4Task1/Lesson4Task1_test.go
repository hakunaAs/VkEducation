package Lesson4Task1

import (
	"testing"
	"time"
)

func TestAlgorithm(t *testing.T) {
	cases := []struct {
		name string
		arr  intArray
		want int
	}{
		{
			"fromExample",
			intArray{95, 87, 100, 92, 85},
			100,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := findMaxNumber(tc.arr)
			if res != tc.want {
				t.Fatal("Неверно найдено максимальное значение!",
					"\nМассив: ", tc.arr,
					"\nФактический результат: ", res, "\nОжидаемый результат: ", tc.want)
			}
		})
	}
}

func startBenchmark(b *testing.B, name string, data intArray, countFunc func(elements intArray) int) {
	b.Run(name, func(b *testing.B) {
		b.ResetTimer()
		start := time.Now()
		for i := 0; i < b.N; i++ {
			_ = countFunc(data)
		}
		elapsed := time.Since(start)
		b.ReportMetric(float64(elapsed.Milliseconds())/float64(b.N), "ms/op")
	})
}

func BenchmarkAlgorithm(b *testing.B) {
	data := make(intArray, 1<<30) // ~1 млд элементов (нужно минимум 16 Гб оперативной памяти)
	for i := range data {
		data[i] = i & 100
	}
	startBenchmark(b, "findMaxNumber", data, findMaxNumber)
}
