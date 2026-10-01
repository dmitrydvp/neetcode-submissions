func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	f := make(map[byte]int8, len(s))

	for i := range len(s) {
		f[s[i]-'a'] += 1

		f[t[i]-'a'] -= 1
	}

	for _, v := range f {
		if v != 0 {
			return false
		}
	}

	return true
}