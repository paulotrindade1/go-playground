package main

import "fmt"

// TwoSum returns the indexes of the list containing the two numbers whose sum is the target
func TwoSum(elements []int, target int) []int {
	complementNumMap := map[int]int{}

	for i, elem := range elements {
		if index, ok := complementNumMap[elem]; ok {
			return []int{index, i}
		}

		complementNumMap[target-elem] = i
	}

	return nil
}

func main() {
	elements := []int{1, 2, 3, 4, 5}

	fmt.Println(TwoSum(elements, 9))
}
