package tracer

import (
	"fmt"
	"io"
)

// Tracer interface
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// New initialize new tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Off creates a tracer that will ignore calls to Trace
func Off() Tracer {
	return &nilTracer{}
}
