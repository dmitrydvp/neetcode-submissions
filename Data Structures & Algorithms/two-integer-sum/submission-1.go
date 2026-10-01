func twoSum(nums []int, target int) []int {
	f := make(map[int]int, len(nums))

	for i, n := range nums {
		difference := target - n

		if _, ok := f[difference]; !ok {
			f[difference] = i
		} else {
			if nums[f[difference]]+n == target {
				return []int{f[difference], i}
			}
		}

		if i == 0 {
			continue
		}

		if nums[f[difference]]+nums[f[target-difference]] == target && f[target-difference] != f[difference] {
			return []int{f[target-difference], f[difference]}
		}
	}

	return []int{}
}