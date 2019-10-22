package rect

// IsOverlap Return true if rc and rc2 is overlapping
func (rc *Rect) IsOverlap(rc2 Rect) bool {
	// check a rect is left of other
	if rc.X > rc2.Right() || rc2.X > rc.Right() {
		return false
	}

	// check a rect is above other
	if rc.Y > rc2.Bottom() || rc2.Y > rc.Bottom() {
		return false
	}

	return true
}

func min(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func (rc *Rect) GetOverlap(rc2 Rect) int16 {
	w := max(0, min(rc.Right(), rc2.Right())-max(rc.X, rc2.X))
	h := max(0, min(rc.Bottom(), rc2.Bottom())-max(rc.Y, rc2.Y))
	return w * h
}
