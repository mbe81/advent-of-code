package mathutil

import (
	"testing"
)

func TestAbs(t *testing.T) {

	testCases := []struct {
		name string
		got  int
		want int
	}{
		{
			name: "positive value",
			got:  1,
			want: 1,
		},
		{
			name: "negative value",
			got:  -1,
			want: 1,
		},
	}

	for _, tc := range testCases {
		if Abs(tc.got) != tc.want {
			t.Errorf("%v: got %d, wanted %d", tc.name, tc.got, tc.want)
		}
	}
}
