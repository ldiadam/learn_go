package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

const numWorkers = 4

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func worker(chunks <-chan []int, wg *sync.WaitGroup, results chan<- []int) {
	defer wg.Done()
	for chunk := range chunks {
		bubbleSort(chunk)
		results <- chunk
	}
}

func mergeSortedSlices(slices [][]int) []int {
	merged := []int{}
	for _, slice := range slices {
		merged = append(merged, slice...)
	}
	sort.Ints(merged) // Sorting the merged slice
	return merged
}

func main() {
	numCount := 1000000
	numbers := make([]int, numCount)
	for i := 0; i < numCount; i++ {
		numbers[i] = rand.Intn(1_000_000) // Random numbers between 0 and 999999
	}

	chunkSize := numCount / numWorkers
	chunks := make(chan []int, numWorkers)
	results := make(chan []int, numWorkers)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(chunks, &wg, results)
	}

	// Send chunks to workers
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = numCount // Handle the last chunk
		}
		chunks <- numbers[start:end]
	}
	close(chunks)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect sorted chunks
	var sortedChunks [][]int
	for sortedChunk := range results {
		sortedChunks = append(sortedChunks, sortedChunk)
	}

	// Merge all sorted chunks
	sortedNumbers := mergeSortedSlices(sortedChunks)

	fmt.Println("Sorting complete. Length of sorted array:", len(sortedNumbers))
}
