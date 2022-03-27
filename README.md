# Aveo

A Environment variables Abstraction System for Go

inspired by [spf13/afero](https://github.com/spf13/afero)

# Go version

go +1.17

# Using Aveo

Aveo is easy to use and easier to adopt.

A few different ways you could use Aveo:

* Wrapper for the OS packages.
* Use Aveo for mock environment variables while testing
* Use the interfaces alone to define your envrioment variables system.
* Define different environment variables for different parts of your application.

## Step 1: Install Aveo

First use go get to install the latest version of the library.

    $ go get github.com/usk81/aveo

Next include Aveo in your application.
```go
import "github.com/usk81/aveo"
```

## Step 2: Declare a backend

First define a package variable and set it to a pointer to a environment variable.
```go
var AppEnv = aveo.NewMap()

or

var AppEnv = aveo.NewOs()
```

## Step 3: Use it like you would the OS package

Throughout your application use any function and method like you normally
would.

So if my application before had:
```go
os.Getenv("foo")
```
We would replace it with:
```go
AppEnv.Getenv("foo")
```

`AppEnv` being the variable we defined above.


## List of all available functions

Environment variable Methods Available:
```go
Clearenv()
Environ() []string
ExpandEnv(s string) string
Getenv(key string) string
LookupEnv(key string) (string, bool)
MapEnvs() map[string]string
Setenv(key, value string) error
Unsetenv(key string) error
```

### How to call

```go
env := aveo.NewMapEnv()
env.Setenv("foo", "bar")
env.Getenv("foo")
```

## Using Aveo for Testing

Before using aveo:
```go
func TestExist(t *testing.T) {
    name := "foo"
    old := os.Getenv(name)

	// set a environment variable
	os.Setenv(name, "bar")
	if got := os.Getenv(name); got == "" {
		t.Errorf("env \"%s\" does not exist.\n", name)
	}
    os.Setenv(name, old)
}
```

Then in my tests I would initialize a new MapEnv for each test:
```go
func TestExist(t *testing.T) {
	env := aveo.NewMapEnv()
	// set a environment variable
    name := "foo"
	env.Setenv(name, "bar")
	if got := env.Getenv(name); got == "" {
		t.Errorf("env \"%s\" does not exist.\n", name)
	}
}
```

# Available Backends

## Operating System Native

### Os

The first is simply a wrapper around the native OS calls. This makes it
very easy to use as all of the calls are the same as the existing OS
calls. It also makes it trivial to have your code use the OS during
operation and a mock environment variable during testing or as needed.

```go
osEnv := aveo.NewOs()
osEnv.Setenv("foo", "bar")
```

## Memory Backed Storage

### Map

Aveo also provides a map data perfect for use in
mocking when persistence isn’t necessary.
It is fully concurrent and will work within go routines safely.

```go
mapEnv := aveo.NewMemMapFs()
mapEnv.Setenv("foo", "bar")
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
