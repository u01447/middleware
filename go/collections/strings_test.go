/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/4 下午5:40
 * Description:
 **/

package collections

import "testing"

func TestMaxNoRepeatedZhn(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"a", 1},
		{"yes, 我爱gogogo", 9},
		{"abcadcb", 4},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		act := MaxNoRepeatedZhn(tt.s)
		if act != tt.ans {
			t.Errorf("get %d for input %s , but expect %d", act, tt.s, tt.ans)
		}
	}
}

func BenchmarkMaxNoRepeatedZhn(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		act := MaxNoRepeatedZhn(s)
		if act != ans {
			b.Errorf("get %d for input %s , but expect %d", act, s, ans)
		}
	}
}
