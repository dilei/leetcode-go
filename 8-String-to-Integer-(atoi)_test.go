package leetcodego

import "testing"

func Test_myAtoi(t *testing.T) {
	str := "  -78"
	expect := -78
	res := myAtoi(str)
	if res != expect {
		t.Error("unexpected")
	}
}
