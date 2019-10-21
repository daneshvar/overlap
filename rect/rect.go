package rect

import "fmt"

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (rc *Rect) Right() int {
	return rc.X + rc.Width
}

func (rc *Rect) Bottom() int {
	return rc.Y + rc.Height
}

func (rc *Rect) ToString() string {
	return fmt.Sprintf("{X: %d, Y: %d, w:%d, h: %d}", rc.X, rc.Y, rc.Width, rc.Height)
}
