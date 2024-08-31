package helpers

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	testcases := []struct {
		in        []string
		want      string
		separator string
	}{
		{[]string{"Hello", "World"}, "Hello World", " "},
		{[]string{" "}, " ", " "},
		{[]string{"divine", "orb"}, "divine orb", " "},
	}

	for _, tc := range testcases {
		result := SliceJoinStrings(tc.in, tc.separator)
		if result != tc.want {
			t.Errorf("Join slices, in: %v want %v separator - %v", tc.in, tc.want, tc.separator)
		}

	}

}
