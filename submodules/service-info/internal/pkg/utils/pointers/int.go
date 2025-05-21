package pointers

// Int returns the integer pointer of an integer.
func Int(integer int) *int {
	return &integer
}
