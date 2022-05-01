package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "Hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Logf("unicode: %x", c[0])
	t.Logf("UTF-8: %x", s)
}

func TestStringtoRune(t *testing.T) {
	s := "海上鋼琴師"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}