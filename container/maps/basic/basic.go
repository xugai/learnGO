package basic


func getRepeatingSubStringLength(str string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0

	// 将字符串转换为字节数组
	for i, ch := range []rune(str) {
		if position, ok := lastOccured[ch]; ok && position >= start {
			start = position + 1
		}
		// 更新不重复子串的长度
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
