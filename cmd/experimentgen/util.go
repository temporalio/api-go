package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func run(workDir string, name string, args ...string) error {
	_, err := runOutput(workDir, name, args...)
	return err
}

func runOutput(workDir string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = workDir
	cmd.Env = append(slices.Clone(os.Environ()), "GOWORK=off")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s %s failed: %w\n%s", name, strings.Join(args, " "), err, out)
	}
	return string(out), nil
}
