package aveo

import (
	"os"
	"strings"
)

// OsFs is a Env implementation that uses functions provided by the os package.
//
// For details in any method, check the documentation of the os package
// (http://golang.org/pkg/os/).
type OsEnv struct{}

func NewOs() Env {
	return &OsEnv{}
}

func (o *OsEnv) Clearenv() {
	os.Clearenv()
}

func (o *OsEnv) Environ() []string {
	return os.Environ()
}

func (o *OsEnv) ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func (o *OsEnv) Getenv(key string) string {
	return os.Getenv(key)
}

func (o *OsEnv) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (o *OsEnv) MapEnvs() map[string]string {
	es := os.Environ()
	m := map[string]string{}
	for _, e := range es {
		kv := strings.SplitN(e, "=", 1)
		m[kv[0]] = kv[1]
	}
	return m
}

func (o *OsEnv) Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func (o *OsEnv) Unsetenv(key string) error {
	return os.Unsetenv(key)
}
