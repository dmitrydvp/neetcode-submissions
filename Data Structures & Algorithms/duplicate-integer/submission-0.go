func hasDuplicate(nums []int) bool {
	f := make(map[int]struct{}, len(nums))

	for _, i := range nums {
		if _, ok := f[i]; !ok {
			f[i] = struct{}{}

			continue
		}

		return true
	}

	return false
}