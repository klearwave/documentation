package pointers

// FromString returns the string pointer of a string.
func FromString(str string) *string {
	return &str
}

// ToString returns the string value of an existing string pointer.
func ToString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

// EqualString returns whether the value of two string pointers are equal.
func EqualString(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}

	if a != nil && b != nil {
		return *a == *b
	}

	return false
}
