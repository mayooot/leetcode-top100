package _3_LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	pre, res := 0, 1
	m := make(map[byte]int, len(s))
	for cur := 0; cur < len(s); cur++ {
		if _, ok := m[s[cur]]; ok {
			pre = max(pre, m[s[cur]]+1)
			m[s[cur]] = cur
			res = max(res, cur-pre+1)
		} else {
			m[s[cur]] = cur
			res = max(res, cur-pre+1)
		}
	}
	return res
}
