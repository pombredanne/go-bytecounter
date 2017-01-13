package bytecounter

import (
	"bytes"
	"io"
	"testing"
)

func TestByteCounter(t *testing.T) {
	buf := &bytes.Buffer{}

	triggered := false
	bc := NewByteCounter(buf, 1, func() {
		triggered = true
	})

	n, err := io.Copy(bc, bytes.NewBuffer([]byte("1")))
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("wrong number of bytes written: %d", n)
	}
	if !triggered {
		t.Fatal("expected action to be triggered")
	}
}

func TestByteCounterOnce(t *testing.T) {
	buf := &bytes.Buffer{}

	triggered := false
	bc := NewByteCounter(buf, 1, func() {
		triggered = true
	})

	io.Copy(bc, bytes.NewBuffer([]byte("1")))
	if !triggered {
		t.Fatal("expected action to be triggered")
	}

	// the action should only trigger once
	triggered = false
	io.Copy(bc, bytes.NewBuffer([]byte("2")))
	if triggered {
		t.Fatal("did not expect action to be triggered")
	}
}
