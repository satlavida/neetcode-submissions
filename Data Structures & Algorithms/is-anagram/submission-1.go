func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ccs := make(map[byte]int, len(s))
	cct := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		cct[s[i]]++
		ccs[t[i]]++
	}

	for k := range ccs {
		if ccs[k] != cct[k] {
			return false
		}
	}
	return true
}
