package Lesson3Task2

import (
	"slices"
	"testing"
	"time"
)

func TestAlgorithm(t *testing.T) {
	cases := []struct {
		name     string
		arr      intArray
		toRemove int
		want     intArray
	}{
		{
			"fromExample",
			intArray{1, 2, -1, 4, 5, -1, 6},
			-1,
			intArray{1, 2, 4, 5, 6},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := DeleteExtraNums(tc.arr, tc.toRemove)
			if !slices.Equal(res, tc.want) {
				t.Fatal("Неверно отсортирован массив!",
					"\nМассив: ", tc.arr,
					"\nФактический результат: ", res, "\nОжидаемый результат: ", tc.want)
			}
		})
	}
}

func startBenchmark(b *testing.B, name string, data intArray, toRemove int, countFunc func(elements intArray, toRemove int) intArray) {
	b.Run(name, func(b *testing.B) {
		b.ResetTimer()
		start := time.Now()
		for i := 0; i < b.N; i++ {
			_ = countFunc(data, toRemove)
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
	startBenchmark(b, "deleteExtraNums", data, 2, DeleteExtraNums)
}
