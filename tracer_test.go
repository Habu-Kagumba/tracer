package tracer

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return from New shouldn't be nil!")
	} else {
		tracer.Trace("Hello tracer package.")
		if buf.String() != "Hello tracer package.\n" {
			t.Errorf("Trace shouldn't write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("Something")
}
