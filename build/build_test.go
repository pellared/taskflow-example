package main

import (
	"strings"
	"testing"

	"github.com/pellared/taskflow"
)

func Test_taskHello(t *testing.T) {
	out := &strings.Builder{}
	r := taskflow.Runner{Output: out}

	got := r.Run(taskHello().Command)

	if got.Failed() {
		t.Errorf("task hello failed")
	}

	if want, gotOutput := "Hello world\n", out.String(); gotOutput != want {
		t.Errorf("got: %v; want: %v", gotOutput, want)
	}
}
