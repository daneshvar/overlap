package main

import (
	"daneshvar/overlap/rect"
	"fmt"
)

func main() {
	rects := [][2]rect.Rect{
		{
			{
				X:      2,
				Y:      1,
				Width:  6,
				Height: 4,
			},
			{
				X:      2,
				Y:      1,
				Width:  6,
				Height: 4,
			},
		},
	}

	for _, r := range rects {
		check(r[0], r[1])
	}
}

func check(rc1 rect.Rect, rc2 rect.Rect) {
	fmt.Printf("%s & %s -> %v\n", rc1.ToString(), rc2.ToString(), rc1.IsOverlap(rc2))
	fmt.Printf("%s & %s -> %d\n", rc1.ToString(), rc2.ToString(), rc1.GetOverlap(rc2))
}
