package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func (g generator) baseSHA(apiRepo string, sourceSHA string) (string, error) {
	for _, ref := range []string{"origin/master", "master", "origin/main", "main"} {
		if _, err := g.runOutput("", "git", "-C", apiRepo, "rev-parse", ref); err != nil {
			continue
		}
		baseSHA, err := g.runOutput("", "git", "-C", apiRepo, "merge-base", sourceSHA, ref)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(baseSHA), nil
	}
	return "", fmt.Errorf("could not determine base SHA for %s", sourceSHA)
}

func (g generator) detectStableRoot(apiRepo string, sourceSHA string) (string, error) {
	for _, candidate := range []string{"temporal/api", ""} {
		path := "workflowservice/v1/request_response.proto"
		if candidate != "" {
			path = filepath.ToSlash(filepath.Join(candidate, path))
		}
		if err := g.run("", "git", "-C", apiRepo, "cat-file", "-e", fmt.Sprintf("%s:%s", sourceSHA, path)); err == nil {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("could not detect stable proto root in %s at %s", apiRepo, sourceSHA)
}

func (g generator) run(workDir string, name string, args ...string) error {
	_, err := g.runOutput(workDir, name, args...)
	return err
}

func (g generator) runOutput(workDir string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = workDir
	cmd.Env = append(slices.Clone(os.Environ()), "GOWORK=off")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s %s failed: %w\n%s", name, strings.Join(args, " "), err, out)
	}
	return string(out), nil
}

var defaultStableVersionResolver = func(sourceSHA string) (moduleVersion, error) {
	module, err := resolveModuleVersion("go.temporal.io/api@" + sourceSHA)
	if err == nil {
		return module, nil
	}

	module, fallbackErr := resolveLocalModuleVersion("go.temporal.io/api")
	if fallbackErr != nil {
		return moduleVersion{}, fmt.Errorf("resolve go.temporal.io/api for source %s failed: %w; fallback failed: %w", sourceSHA, err, fallbackErr)
	}
	return module, nil
}

func resolveModuleVersion(path string) (moduleVersion, error) {
	out, err := exec.Command("go", "list", "-m", "-json", path).CombinedOutput()
	if err != nil {
		return moduleVersion{}, fmt.Errorf("go list -m %s failed: %w\n%s", path, err, out)
	}
	var result struct {
		Version   string
		GoVersion string
	}
	if err := json.Unmarshal(out, &result); err != nil {
		return moduleVersion{}, err
	}
	if result.Version == "" || result.GoVersion == "" {
		return moduleVersion{}, fmt.Errorf("could not resolve module version for %s", path)
	}
	return moduleVersion{
		Version:   result.Version,
		GoVersion: result.GoVersion,
	}, nil
}

func resolveLocalModuleVersion(path string) (moduleVersion, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return moduleVersion{}, fmt.Errorf("could not locate generator source root")
	}
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	goModPath := filepath.Join(repoRoot, "go.mod")
	goModBytes, err := os.ReadFile(goModPath)
	if err != nil {
		return moduleVersion{}, err
	}
	goVersion := ""
	for _, line := range strings.Split(string(goModBytes), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "go ") {
			goVersion = strings.TrimSpace(strings.TrimPrefix(line, "go "))
			break
		}
	}
	if goVersion == "" {
		return moduleVersion{}, fmt.Errorf("could not determine go version from %s", goModPath)
	}

	cmd := exec.Command("git", "-C", repoRoot, "describe", "--tags", "--match", "v*", "--abbrev=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return moduleVersion{}, fmt.Errorf("git describe for %s failed: %w\n%s", path, err, out)
	}
	return moduleVersion{
		Version:   strings.TrimSpace(string(out)),
		GoVersion: goVersion,
	}, nil
}
