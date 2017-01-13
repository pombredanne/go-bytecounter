package bytecounter

import (
	"io"
	"sync"
)

// ByteCounter counts the bytes passing through an io.Writer
// and triggers an action when a certain threshold is passed
// Note the trigger function is only called once
type ByteCounter struct {
	w     io.Writer
	count int64
	fn    func()
	once  *sync.Once
}

// NewByteCounter wraps an existing io.Writer with a ByteCounter
func NewByteCounter(w io.Writer, count int64, fn func()) io.Writer {
	return &ByteCounter{
		w:     w,
		count: count,
		fn:    fn,
		once:  &sync.Once{},
	}
}

func (w *ByteCounter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.count -= int64(n)
	if w.count <= 0 {
		w.once.Do(func() {
			w.fn()
		})
	}
	return n, err
}
