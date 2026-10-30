package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProtocPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("test executables use shell scripts")
	}

	dir := t.TempDir()
	protogen := filepath.Join(dir, "protogen")
	output, err := exec.Command("go", "build", "-o", protogen, ".").CombinedOutput()
	require.NoError(t, err, "%s", output)
	protoRoot := filepath.Join(dir, "protos")
	require.NoError(t, os.Mkdir(protoRoot, 0755))
	protoFile := filepath.Join(protoRoot, "test.proto")
	require.NoError(t, os.WriteFile(protoFile, []byte("syntax = \"proto3\";\n"), 0644))
	binDir := filepath.Join(dir, "bin")
	require.NoError(t, os.Mkdir(binDir, 0755))
	defaultProtoc := filepath.Join(binDir, "protoc")
	customProtoc := filepath.Join(dir, "custom protoc")
	for _, path := range []string{defaultProtoc, customProtoc} {
		require.NoError(t, os.WriteFile(path, []byte("#!/bin/sh\nprintf '%s\\n' \"$0\" \"$@\" > \"$PROTOGEN_TEST_ARGS_FILE\"\n"), 0755))
	}

	for _, tc := range []struct {
		name       string
		protocPath string
		wantPath   string
		wantError  bool
	}{
		{name: "default", wantPath: defaultProtoc},
		{name: "explicit", protocPath: customProtoc, wantPath: customProtoc},
		{name: "missing", protocPath: filepath.Join(dir, "missing-protoc"), wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			argsFile := filepath.Join(t.TempDir(), "args")
			descriptor := filepath.Join(dir, "descriptor.pb")
			args := []string{"--root=" + protoRoot, "--output-descriptor=" + descriptor, "--no-rewrite-enum-const", "--no-rewrite-enum-string", "--no-strip-version"}
			if tc.protocPath != "" {
				args = append(args, "--protoc-path="+tc.protocPath)
			}
			cmd := exec.Command(protogen, args...)
			cmd.Env = append(os.Environ(), "PATH="+binDir, "PROTOGEN_TEST_ARGS_FILE="+argsFile)
			output, err := cmd.CombinedOutput()
			if tc.wantError {
				require.Error(t, err)
				require.Contains(t, string(output), tc.protocPath)
				require.NoFileExists(t, argsFile)
				return
			}
			require.NoError(t, err, "%s", output)
			got, err := os.ReadFile(argsFile)
			require.NoError(t, err)
			require.Equal(t, tc.wantPath+"\n--fatal_warnings\n--include_imports\n--include_source_info\n--descriptor_set_out="+descriptor+"\n-I="+protoRoot+"\n"+protoFile+"\n", string(got))
		})
	}
}
