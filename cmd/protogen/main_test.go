package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestCompileProtosExclusions(t *testing.T) {
	root := t.TempDir()
	files := []string{
		"excluded/nested/skip.proto",
		"excluded/skip.proto",
		"excluded_other/keep.proto",
		"keep.proto",
	}
	for _, file := range files {
		path := filepath.Join(root, filepath.FromSlash(file))
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0755))
		require.NoError(t, os.WriteFile(path, []byte("syntax = \"proto3\";\n"), 0644))
	}

	for _, tt := range []struct {
		name      string
		exclusion string
		wantFiles []string
		wantError string
	}{
		{
			name:      "directory",
			exclusion: "excluded",
			wantFiles: []string{"excluded_other/keep.proto", "keep.proto"},
		},
		{
			name:      "noncanonical directory path",
			exclusion: "./excluded/",
			wantFiles: []string{"excluded_other/keep.proto", "keep.proto"},
		},
		{
			name:      "nested directory",
			exclusion: "excluded/nested",
			wantFiles: []string{"excluded/skip.proto", "excluded_other/keep.proto", "keep.proto"},
		},
		{
			name:      "missing directory",
			exclusion: "missing",
			wantFiles: files,
		},
		{
			name:      "file",
			exclusion: "keep.proto",
			wantError: `--exclude "keep.proto" must name a directory`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			descriptorPath := filepath.Join(t.TempDir(), "descriptor.pb")
			err := compileProtos(t.Context(), genConfig{
				protocPath:           "protoc",
				rootDirs:             []string{root},
				excludeDirs:          []string{filepath.FromSlash(tt.exclusion)},
				includes:             []string{root},
				outputDescriptorPath: descriptorPath,
			})
			if tt.wantError != "" {
				require.ErrorContains(t, err, tt.wantError)
				return
			}
			require.NoError(t, err)

			data, err := os.ReadFile(descriptorPath)
			require.NoError(t, err)
			var descriptors descriptorpb.FileDescriptorSet
			require.NoError(t, proto.Unmarshal(data, &descriptors))

			var names []string
			for _, file := range descriptors.GetFile() {
				names = append(names, file.GetName())
			}
			require.Equal(t, tt.wantFiles, names)
		})
	}
}
