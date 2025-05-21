package types

import (
	"strings"
)

// ToBoolean converts an interface to its boolean value.  It return false if the type is not known
// or cannot be converted.
//
//nolint:forcetypeassert // since we are switching we can safely assume the type assertion works
func ToBoolean(in any) bool {
	switch t := in.(type) {
	case string:
		for _, boolVal := range []string{
			"true",
			"t",
			"yes",
		} {
			if strings.EqualFold(t, boolVal) {
				return true
			}
		}
	case int, int8, int16, int32, int64:
		val := t.(int)

		return val == 1
	case bool:
		return t
	}

	return false
}
