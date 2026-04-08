package lsubstring

func LSubString(line string) int {
	left := 0
	maxLen := 0
	currLen := 0

	m := make(map[byte]int)
	if len(line) == 1 {
		return 1
	}
	for right := 0; right < len(line); right++ {
		char := line[right]
		if idx, ok := m[char]; ok && idx >= left {
			left = idx + 1
		}
		m[char] = right
		currLen = right - left + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}
