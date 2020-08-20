package domain

import "fmt"

// ConflictStrategy define how a conflict of an already existing branch should be handled
type ConflictStrategy int

const (
	// ConflictStrategySkip will skip the run for if the branch does already exist
	ConflictStrategySkip ConflictStrategy = iota + 1
	// ConflictStrategyAppend will run the script on top of the already existing branch
	ConflictStrategyAppend
	// ConflictStrategyReplace will ignore any existing branch and replace it with new changes
	ConflictStrategyReplace
)

func ParseConflictStrategy(str string) (ConflictStrategy, error) {
	switch str {
	default:
		return ConflictStrategy(0), fmt.Errorf("could not parse \"%s\" as conflict strategy", str)
	case "skip":
		return ConflictStrategySkip, nil
	case "append":
		return ConflictStrategyAppend, nil
	case "replace":
		return ConflictStrategyReplace, nil
	}
}
