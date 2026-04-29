package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func newExampleAPIRepo(t *testing.T) string {
	t.Helper()

	apiRepo := t.TempDir()
	runCmd(t, "", "git", "init", apiRepo)
	runCmd(t, apiRepo, "git", "config", "user.email", "test@example.com")
	runCmd(t, apiRepo, "git", "config", "user.name", "Test User")

	copyFixtureTree(t, filepath.Join("testdata", "example", "base"), apiRepo)

	runCmd(t, apiRepo, "git", "add", ".")
	runCmd(t, apiRepo, "git", "commit", "-m", "base")

	runCmd(t, apiRepo, "git", "checkout", "-b", "example")
	copyFixtureTree(t, filepath.Join("testdata", "example", "feature"), apiRepo)
	runCmd(t, apiRepo, "git", "add", ".")
	runCmd(t, apiRepo, "git", "commit", "-m", "add example proto")
	return apiRepo
}

func copyFixtureTree(t *testing.T, srcRoot string, dstRoot string) {
	t.Helper()

	entries, err := os.ReadDir(srcRoot)
	require.NoError(t, err)
	for _, entry := range entries {
		srcPath := filepath.Join(srcRoot, entry.Name())
		dstPath := filepath.Join(dstRoot, entry.Name())
		if entry.IsDir() {
			require.NoError(t, os.MkdirAll(dstPath, 0o755))
			copyFixtureTree(t, srcPath, dstPath)
			continue
		}
		contents, err := os.ReadFile(srcPath)
		require.NoError(t, err)
		require.NoError(t, os.MkdirAll(filepath.Dir(dstPath), 0o755))
		require.NoError(t, os.WriteFile(dstPath, contents, 0o644))
	}
}

func runCmd(t *testing.T, dir string, name string, args ...string) string {
	t.Helper()
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, string(out))
	return string(out)
}
