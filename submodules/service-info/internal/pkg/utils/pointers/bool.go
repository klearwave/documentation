package pointers

// Bool returns the boolean pointer of a boolean.
func Bool(boolean bool) *bool {
	return &boolean
}
