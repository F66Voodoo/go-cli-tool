package service

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()

	return buf.String()
}

func TestPrintStartMessage_NoName(t *testing.T) {
	output := captureStdout(t, func() {
		PrintStartMessage("")
	})

	if !strings.Contains(output, "go-cli-tool started") {
		t.Fatalf("unexpected output: %q", output)
	}
}

func TestPrintStartMessage_WithName(t *testing.T) {
	output := captureStdout(t, func() {
		PrintStartMessage("Pavel")
	})

	if !strings.Contains(output, "Pavel") {
		t.Fatalf("unexpected output: %q", output)
	}
}
