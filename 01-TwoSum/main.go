package _1_TwoSum

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := set[target-nums[i]]; ok {
			return []int{set[target-nums[i]], i}
		}
		set[nums[i]] = i
	}
	return nil
}
