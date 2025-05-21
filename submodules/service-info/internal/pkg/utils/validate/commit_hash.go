package validate

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidCommitHash = errors.New("invalid commit hash value")
	ErrMissingCommitHash = errors.New("missing commit hash value")
)

const (
	commitHashRegex = `^[a-fA-F0-9]{40}$`
)

// CommitHash checks if the given Git commit hash is valid.
func CommitHash(commitHash string) error {
	if commitHash == "" {
		return ErrMissingCommitHash
	}

	re, err := regexp.Compile(commitHashRegex)
	if err != nil {
		return fmt.Errorf("invalid commit hash regex pattern [%s]: %w", commitHashRegex, err)
	}

	if !re.MatchString(commitHash) {
		return fmt.Errorf("commit hash [%s] does not match pattern [%s]; %w",
			commitHash,
			commitHashRegex,
			ErrInvalidCommitHash,
		)
	}

	return nil
}
