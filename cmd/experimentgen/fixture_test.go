package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func buildTestDescriptorSet(t *testing.T) []byte {
	t.Helper()
	protoRoot := filepath.Join("testdata", "example", "protos")
	cmd := exec.Command("buf", "build", "-o", "-")
	cmd.Dir = protoRoot
	out, err := cmd.Output()
	if err != nil {
		if _, lookErr := exec.LookPath("buf"); lookErr != nil {
			t.Skip("buf not available")
		}
		t.Fatalf("buf build: %v: %s", err, out)
	}
	return out
}

func readFile(t *testing.T, dir string, rel string) string {
	t.Helper()
	contents, err := os.ReadFile(filepath.Join(dir, rel))
	if err != nil {
		t.Fatalf("readFile %s: %v", rel, err)
	}
	return string(contents)
}
