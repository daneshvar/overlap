package rect

import "fmt"

type Rect struct {
	X      int16
	Y      int16
	Width  int16
	Height int16
}

func (rc *Rect) Right() int16 {
	return rc.X + rc.Width
}

func (rc *Rect) Bottom() int16 {
	return rc.Y + rc.Height
}

func (rc *Rect) ToString() string {
	return fmt.Sprintf("{X: %d, Y: %d, w:%d, h: %d}", rc.X, rc.Y, rc.Width, rc.Height)
}
