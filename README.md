stacktrace
==========

A Go library to get stacktrace for logging using only one shared buffer without other buffer memory allocation.
Note memory allocations may occur in [runtime.Stack](http://golang.org/pkg/runtime/#Stack).

## Usage

First, define a function to print error message with the stacktrace for your favorite logging library.
This library uses the one shared buffer, so you have to unlock it after logging the stacktrace.

```
func logErrorWithStackTrace(msg interface{}) {
	log.Printf("error: %s\nstacktrace: %s\n", msg, stacktrace.LockBufAndGetStack())
	defer stacktrace.UnlockBuf()
}
```

Then use it.

```
func b() {
	err := errors.New("some error")
	logErrorWithStackTrace(err)
}
```

See the example/main.go for the full source code.

## License
MIT
