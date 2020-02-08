package basic

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{a, b, c int}{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		//{30000, 40000, 50000},
	}

	for _, test := range tests {
		if actual := calcTriangle(test.a, test.b); actual != test.c {
			t.Errorf("calcTriangle(%d, %d) is incorrect, actual %d, excepted %d\n",
					test.a, test.b, actual, test.c)
		}
	}
}