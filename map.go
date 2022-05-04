package aveo

import (
	"os"
	"syscall"
)

// MapEnv is designed for use in testing.
// Not recommended for production use
type MapEnv struct {
	store map[string]string
}

func NewMap(store map[string]string) Env {
	if store == nil {
		store = make(map[string]string)
	}
	return &MapEnv{
		store: store,
	}
}

func (m *MapEnv) Clearenv() {
	m.store = make(map[string]string)
}

func (m *MapEnv) Environ() []string {
	em := m.MapEnvs()
	envs := make([]string, len(em))
	var i int
	for k, v := range em {
		envs[i] = k + "=" + v
		i++
	}
	return envs
}

func (m *MapEnv) ExpandEnv(s string) string {
	return os.Expand(s, m.Getenv)
}

func (m *MapEnv) Getenv(key string) string {
	if v, ok := m.LookupEnv(key); ok {
		return v
	}
	return ""
}

func (m *MapEnv) LookupEnv(key string) (string, bool) {
	v, ok := m.store[key]
	return v, ok
}

func (m *MapEnv) MapEnvs() map[string]string {
	return m.store
}

func (m *MapEnv) Setenv(key, value string) error {
	if key == "" {
		return os.NewSyscallError("setenv", syscall.EINVAL)
	}
	if value != "" {
		m.store[key] = value
	}
	return nil
}

func (m *MapEnv) Unsetenv(key string) error {
	delete(m.store, key)
	return nil
}
