package main

import "fmt"

func chunkRecords(records []int, chunkSize int) [][]int {
	var recordsBulk [][]int

	for i := 0; i < len(records)/chunkSize; i++ {
		recordsBulk = append(recordsBulk, records[:chunkSize])
		records = records[chunkSize:]
	}

	chunkMod := len(records) % chunkSize
	if chunkMod > 0 {
		recordsBulk = append(recordsBulk, records[:chunkMod])
	}

	return recordsBulk
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(chunkRecords(ints, 3))
}
