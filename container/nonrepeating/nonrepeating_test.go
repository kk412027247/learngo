package main

import "testing"

func TestSubstring (t *testing.T) {
	tests := []struct{
		s string
		ans int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef",6},
		{"一二三二一一二三二一一二三二一一二三二一一二三二一", 3},
	}

	for _, tt:= range tests {
		actual :=  lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for intpu %s; expected %d", actual, tt.s, tt.ans)
		}
	}

}

// 性能测试
func BenchmarkSubstring(b *testing.B){
	s := "一二三二一一二三二一一二三二一一二三二一一二三二一"

	for i := 0; i<13;i++{
		s = s + s
	}
	b.Logf("len(s)=%d", len(s))

	ans := 3
  	// 以上的运行不计入测试时间
	b.ResetTimer()
	for i :=0 ; i < b.N ; i++ {
		actual :=  lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for intpu %s; expected %d", actual, s, ans)
		}
	}



}
