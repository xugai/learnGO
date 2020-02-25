package basic

// 用空间换时间
var lastOccured = make([]int, 0xffff)

func getRepeatingSubStringLength(str string) int {

	//lastOccured := make(map[rune]int)
	for index := range lastOccured {
		lastOccured[index] = -1
	}
	start := 0
	maxLength := 0


	// 将字符串转换为字节数组
	for i, ch := range []rune(str) {
		if position := lastOccured[ch]; position != -1 && position >= start {
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
