func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	f := make(map[byte]struct {
		s uint8
		t uint8
	}, len(s))

	for i := range len(s) {
		if _, ok := f[s[i]]; !ok {
			f[s[i]] = struct {
				s uint8
				t uint8
			}{s: 1, t: 0}
		} else {
			f[s[i]] = struct {
				s uint8
				t uint8
			}{s: f[s[i]].s + 1, t: f[s[i]].t}
		}

		if _, ok := f[t[i]]; !ok {
			f[t[i]] = struct {
				s uint8
				t uint8
			}{s: 0, t: 1}
		} else {
			f[t[i]] = struct {
				s uint8
				t uint8
			}{s: f[t[i]].s, t: f[t[i]].t + 1}
		}
	}

	for _, v := range f {
		if v.s != v.t {
			return false
		}
	}

	return true
}