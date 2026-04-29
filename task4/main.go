package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // เก็บค่า: ตัวเลข -> index

	for i := 0; i < len(nums); i++ {
		need := target - nums[i]

		// เช็คว่ามีตัวที่ต้องการอยู่ก่อนหน้านี้ไหม
		if index, ok := m[need]; ok {
			return []int{index, i}
		}

		// ยังไม่เจอ → เก็บค่าปัจจุบันไว้ใน map
		m[nums[i]] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println("Result:", result)
}