package main

import (
	"fmt"
	//"unicode/utf8"
	"strings"
)

func main() {

	//fmt.Println("Creating map:")
	//m := map[string]string {
	//	"name": "berio xu",
	//	"course": "golang",
	//	"site": "imooc",
	//	"quality": "good",
	//}		// 直接赋初值的方式创建map
	//
	//fmt.Println("Traversing map:")
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	//
	//m2 := make(map[string]int)		// m2 == empty map 通过调用make函数的方式创建map
	//var m3 map[string]int		// m3 == nil
	//fmt.Println(m, m2, m3)
	//
	//fmt.Println("Getting values:")
	//if name, ok := m["nname"]; ok {
	//	fmt.Println(name)
	//} else {
	//	fmt.Println("The key you want to find doesn't exist!")
	//}
	//delete(m, "name")
	//name, ok := m["name"]
	//fmt.Println(name, ok)
	//
	//fmt.Println("Restore value")
	//m["name"] = "Behe"
	//name, ok = m["name"]
	//fmt.Println(name, ok)

	fmt.Println(getRepeatingSubStringLength("abcabcab"))
	fmt.Println(getRepeatingSubStringLength("aaaaa"))
	fmt.Println(getRepeatingSubStringLength("abcdef"))
	fmt.Println(getRepeatingSubStringLength("你好徐秋冰徐改"))
	fmt.Println(getRepeatingSubStringLength("我不想想不我"))

	fmt.Println(strings.Fields("a  b    c   d"))

	//str := "Yes我爱慕课网"  // UTF-8编码，英文字母占一个字节，中文占三个字节
	//for i, ch := range str {  // ch is a rune
	//	// 将str进行UTF-8解码，解出来后又将每一个字符进行unicode编码，转完后将值存放在rune里
	//	fmt.Printf("(%d, %X) ", i, ch)
	//}
	//fmt.Println()
	//for i, byte := range []byte(str) {
	//	fmt.Printf("(%d, %X) ", i, byte)
	//}
	//fmt.Println()
	//for i, ch := range []rune(str) {
	//	fmt.Printf("(%d, %c) ", i, ch)
	//}
	//fmt.Println()
	//fmt.Println("Rune Count: ", utf8.RuneCountInString(str))
	//
	//bytes := []byte(str)
	//for len(bytes) > 0 {
	//	ch, size := utf8.DecodeRune(bytes)
	//	bytes = bytes[size:]
	//	fmt.Printf("%c", ch)
	//}
	//fmt.Println()
}

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
