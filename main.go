package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	maxNum := data[0]
	for _, num := range data {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	res := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		sliceBegin := i * (SIZE / CHUNKS)
		sliceEnd := sliceBegin + (SIZE / CHUNKS)
		if i == CHUNKS-1 {
			sliceEnd = SIZE
		}
		sliceOfData := data[sliceBegin:sliceEnd]
		go func(id int, dataCh []int) {
			defer wg.Done()
			max := maximum(sliceOfData)
			res[id] = max
		}(i, sliceOfData)
	}
	wg.Wait()
	return maximum(res)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
