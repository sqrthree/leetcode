package main

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		v := target - nums[i]

		if _, ok := m[v]; ok {
			return []int{m[v], i}
		}

		m[nums[i]] = i
	}

	return nil
}
