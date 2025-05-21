package validate

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidSHA256Sum = errors.New("invalid sha256sum value")
	ErrMissingSHA256Sum = errors.New("missing sha256sum value")
)

const (
	sha256Regex = `^[a-fA-F0-9]{64}$`
)

// SHA256Sum checks if the given SHA256 sum is valid.
func SHA256Sum(sha256sum string) error {
	if sha256sum == "" {
		return ErrMissingSHA256Sum
	}

	re, err := regexp.Compile(sha256Regex)
	if err != nil {
		return fmt.Errorf("invalid sha256sum regex pattern [%s]: %w", sha256Regex, err)
	}

	if !re.MatchString(sha256sum) {
		return fmt.Errorf("sha256sum [%s] does not match pattern [%s]; %w",
			sha256sum,
			sha256Regex,
			ErrInvalidSHA256Sum,
		)
	}

	return nil
}
