package rect

import "testing"

var tests = []struct {
	rc1     Rect
	rc2     Rect
	overlap int
}{
	{
		rc1:     Rect{2, 1, 6, 4},
		rc2:     Rect{2, 1, 6, 4},
		overlap: 24,
	},
}

func TestIsOverlap(t *testing.T) {
	for i, test := range tests {
		in := test.overlap > 0
		out := test.rc1.IsOverlap(test.rc2)
		if out != in {
			t.Errorf("#%d: %s & %s -> %v; want %v", i, test.rc1.ToString(), test.rc2.ToString(), out, in)
		}
	}
}

func TestGetOverlap(t *testing.T) {
	for i, test := range tests {
		in := test.overlap
		out := test.rc1.GetOverlap(test.rc2)
		if out != in {
			t.Errorf("#%d: %s & %s -> %d; want %d", i, test.rc1.ToString(), test.rc2.ToString(), out, in)
		}
	}
}
