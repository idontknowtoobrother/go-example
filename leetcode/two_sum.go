package leetcode

func TwoSum(nums []int, target int) [2]int {

	tsm := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		idx, found := tsm[diff]

		if found {
			return [2]int{idx, i}
		}

		tsm[num] = i
	}

	return [2]int{-1, -1}
}
