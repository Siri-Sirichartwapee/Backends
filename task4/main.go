package main

import "fmt"

// TwoSum finds indices of two numbers that add up to target.
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int) // value -> index
	for i, num := range nums {
		complement := target - num
		if index, ok := m[complement]; ok {
			return []int{index, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	fmt.Println("--- Task 4: Two Sum ---")
	nums := []int{2, 7, 11, 15}
	target := 9

	indices := TwoSum(nums, target)
	fmt.Printf("Nums: %v, Target: %d\n", nums, target)
	if indices != nil {
		fmt.Printf("Indices: %v (Values: %d + %d = %d)\n", 
			indices, nums[indices[0]], nums[indices[1]], target)
	} else {
		fmt.Println("No solution found.")
	}
}
