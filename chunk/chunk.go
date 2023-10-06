package main

import "fmt"

func chunkRecords(records []int, chunkSize int) [][]int {
	var recordsBulk [][]int

	if chunkSize < 0 {
		return recordsBulk
	}

	for i := 0; i < len(records); i += chunkSize {
		end := i + chunkSize
		if end > len(records) {
			end = len(records)
		}

		recordsBulk = append(recordsBulk, records[i:end])
	}

	return recordsBulk
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(chunkRecords(ints, 3))
}
