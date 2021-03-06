package basic

import "testing"

// golang中的表格驱动测试
// 结构体里面的属性列表之间是不需要加逗号分隔开的
func TestMaps(t *testing.T) {
	tests := [] struct{
		str string
		answer int
	}{
		// Normal case
		{"abcabcabc", 3},
		{"abcdef", 6},

		// Edge case
		{"aaaa", 1},
		{"abcabcabcd", 4},

		// Chinese case
		{"你好徐秋冰徐改", 5},
		{"我不想想不我", 3},
	}


	for _, test := range tests {
		actual := getRepeatingSubStringLength(test.str)
		if actual != test.answer {
			t.Errorf("getRepeatingSubStringLength(%v) is incorrect, actual %d, excepted %d\n",
					test.str, actual, test.answer)
		}
	}

}

func BenchmarkFunc(b *testing.B) {
	str := "你好徐秋冰徐改"

	for i := 0; i < 13; i++ {
		str = str + str
	}

	answer := 6
	for i := 0; i < b.N; i++ {
		if actual := getRepeatingSubStringLength(str); actual != answer {
			b.Errorf("getRepeatingSubStringLength(%v) is incorrect, actual is %d, excepted is %d\n",
				str, actual, answer)
		}
	}


}
