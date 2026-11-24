package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestWalkExtensionExclusions(t *testing.T) {
	tests := []struct {
		name      string
		exclusion string
		files     []string
		want      []string
	}{
		{
			name:      "missing directory",
			exclusion: "temporal/api_next",
			files:     []string{"temporal/api/message.proto"},
			want:      []string{"temporal/api/message.proto"},
		},
		{
			name:      "prefix within a component",
			exclusion: "temporal/api_next",
			files:     []string{"temporal/api_next_other/message.proto"},
			want:      []string{"temporal/api_next_other/message.proto"},
		},
		{
			name:      "nested exclusion",
			exclusion: "temporal/api_next/internal",
			files:     []string{"temporal/api_next/internal/v1/message.proto", "temporal/api_next/public/v1/message.proto"},
			want:      []string{"temporal/api_next/public/v1/message.proto"},
		},
		{
			name:      "excluded directory and sibling",
			exclusion: "temporal/api_next",
			files:     []string{"temporal/api_next/message.proto", "temporal/api/message.proto"},
			want:      []string{"temporal/api/message.proto"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			for _, file := range tt.files {
				path := filepath.Join(root, filepath.FromSlash(file))
				require.NoError(t, os.MkdirAll(filepath.Dir(path), 0755))
				require.NoError(t, os.WriteFile(path, nil, 0644))
			}

			var got []string
			err := walkExtension(t.Context(), root, ".proto", []string{filepath.FromSlash(tt.exclusion)}, func(path string) error {
				rel, err := filepath.Rel(root, path)
				if err != nil {
					return err
				}
				got = append(got, filepath.ToSlash(rel))
				return nil
			})
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCompileProtosExcludesDirectory(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"excluded/skip.proto", "excluded_other/keep.proto"} {
		path := filepath.Join(root, filepath.FromSlash(name))
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0755))
		require.NoError(t, os.WriteFile(path, []byte("syntax = \"proto3\";\n"), 0644))
	}

	descriptorPath := filepath.Join(t.TempDir(), "descriptor.pb")
	require.NoError(t, compileProtos(t.Context(), genConfig{
		protocPath:           "protoc",
		rootDirs:             []string{root},
		excludeDirs:          []string{filepath.FromSlash("./excluded/")},
		includes:             []string{root},
		outputDescriptorPath: descriptorPath,
	}))

	data, err := os.ReadFile(descriptorPath)
	require.NoError(t, err)
	var descriptors descriptorpb.FileDescriptorSet
	require.NoError(t, proto.Unmarshal(data, &descriptors))

	var names []string
	for _, file := range descriptors.GetFile() {
		names = append(names, file.GetName())
	}
	require.Equal(t, []string{"excluded_other/keep.proto"}, names)
}

func TestWalkExtensionRejectsFileExclusion(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(root, "options.proto"), nil, 0644))

	err := walkExtension(t.Context(), root, ".proto", []string{"options.proto"}, func(string) error {
		return nil
	})

	require.ErrorContains(t, err, `--exclude "options.proto" must name a directory`)
}
