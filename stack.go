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

func LockBufAndGetStackWithSkip(skip uint) []byte {
	stackBuf.mu.Lock()

	n := runtime.Stack(stackBuf.buf[:], false)

	start1 := indexLineStart(stackBuf.buf[:n], 1)
	start2 := indexLineStart(stackBuf.buf[:n], 2*skip+1)

	copy(stackBuf.buf[start1:], stackBuf.buf[start2:n])
	return stackBuf.buf[:n-(start2-start1)]
}

func UnlockBuf() {
	stackBuf.mu.Unlock()
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
