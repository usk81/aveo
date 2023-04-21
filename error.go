package aveo

import "errors"

const (
	ErrInvalidEnvvarName  = errors.New("invalid environment variable name")
	ErrInvalidMapItem     = errors.New("invalid map item")
	ErrLookuperNil        = errors.New("lookuper cannot be nil")
	ErrMissingKey         = errors.New("missing key")
	ErrMissingRequired    = errors.New("missing required value")
	ErrNoInitNotPtr       = errors.New("field must be a pointer to have noinit")
	ErrNotPtr             = errors.New("input must be a pointer")
	ErrNotStruct          = errors.New("input must be a struct")
	ErrPrefixNotStruct    = errors.New("prefix is only valid on struct types")
	ErrPrivateField       = errors.New("cannot parse private fields")
	ErrRequiredAndDefault = errors.New("field cannot be required and have a default value")
	ErrUnknownOption      = errors.New("unknown option")
)
