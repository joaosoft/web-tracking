package validator

import "github.com/joaosoft/errors"

var (
	ErrorSkipValidation     = errors.New(errors.LevelError, 1, "skip validation")
	ErrorInvalidValue       = errors.New(errors.LevelError, 2, "invalid value")
	ErrorInvalidPointer     = errors.New(errors.LevelError, 3, "invalid pointer")
	ErrorInvalidTag         = errors.New(errors.LevelError, 4, "invalid tag [%s]")
	ErrorInvalidTagArgument = errors.New(errors.LevelError, 5, "invalid tag argument [%s]")
	ErrorInvalidTagPrefix   = errors.New(errors.LevelError, 6, "invalid prefix [%s] on tag [%s]")
)
