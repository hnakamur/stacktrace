package stacktrace

import (
	"bytes"
	"runtime"
	"sync"
)

// NOTE: You may need to make this larger if your call stacks are deep.
const bufSize = 8192

var stackBuf struct {
	buf [bufSize]byte
	mu  sync.Mutex
}

const lf = 0x0A

func Stack() []byte {
	// NOTE: 1 for this function
	return StackWithSkip(1 + 1)
}

func StackWithSkip(skip uint) []byte {
	stackBuf.mu.Lock()
	defer stackBuf.mu.Unlock()

	n := runtime.Stack(stackBuf.buf[:], false)

	start1 := indexLineStart(stackBuf.buf[:n], 1)
	start2 := indexLineStart(stackBuf.buf[:n], 2*skip+1)

	buf := make([]byte, n-(start2-start1))
	copy(buf, stackBuf.buf[:start1])
	copy(buf[start1:], stackBuf.buf[start2:])
	return buf
}

func indexLineStart(buf []byte, count uint) int {
	offset := 0
	start := 0
	for i := uint(0); i < count; i++ {
		offset += start
		start = bytes.IndexByte(buf[offset:], lf)
		if start == -1 {
			return -1
		}
		start++
	}
	return offset + start
}
