package aveo

type (
	Env interface {
		// Clearenv deletes all environment variables.
		Clearenv()

		// Environ returns a copy of strings representing the environment,
		// in the form "key=value".
		Environ() []string

		// ExpandEnv replaces ${var} or $var in the string according to the values
		// of the current environment variables. References to undefined
		// variables are replaced by the empty string.
		ExpandEnv(s string) string

		// Getenv retrieves the value of the environment variable named by the key.
		// It returns the value, which will be empty if the variable is not present.
		// To distinguish between an empty value and an unset value, use LookupEnv.
		Getenv(key string) string

		// LookupEnv retrieves the value of the environment variable named
		// by the key. If the variable is present in the environment the
		// value (which may be empty) is returned and the boolean is true.
		// Otherwise the returned value will be empty and the boolean will
		// be false.
		LookupEnv(key string) (string, bool)

		// Environ returns a copy of strings representing the environment,
		// as map[key]value.
		MapEnvs() map[string]string

		// Setenv sets the value of the environment variable named by the key.
		// It returns an error, if any.
		Setenv(key, value string) error

		// Unsetenv unsets a single environment variable.
		Unsetenv(key string) error
	}
)
