package trace_test

import (
	"bytes"
	"testing"

	"github.com/l0s0s/WebSocketChat/trace"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	if tracer := trace.New(&buf); tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer trace.Tracer = trace.Off()
	silentTracer.Trace("something")
}
