package leetcode

func twoSum(nums []int, target int) []int {
	numslength := len(nums)
	for i, v := range nums {
		for j := i + 1; j < numslength; j++ {
			if nums[j] == target-v {
				return []int{i, j}
			}
		}
	}
	return nil
}