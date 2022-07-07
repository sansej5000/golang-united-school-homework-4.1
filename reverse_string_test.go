package reverse_string

import "testing"

func Test_ReverseString(t *testing.T) {
	expected := "esrever"
	input := "reverse"
	res := ReverseString(input)
	if expected != res {
		t.Errorf("invalid value, expected %s", expected)
	}
}
