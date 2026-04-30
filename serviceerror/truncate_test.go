package serviceerror

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func TestTruncateUTF8(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		n      int
		expect string
	}{
		{
			name:   "empty string",
			input:  "",
			n:      10,
			expect: "",
		},
		{
			name:   "n is zero",
			input:  "hello",
			n:      0,
			expect: "",
		},
		{
			name:   "n is negative",
			input:  "hello",
			n:      -1,
			expect: "",
		},
		{
			name:   "ascii shorter than n",
			input:  "hello",
			n:      10,
			expect: "hello",
		},
		{
			name:   "ascii equal to n",
			input:  "hello",
			n:      5,
			expect: "hello",
		},
		{
			name:   "ascii longer than n",
			input:  "hello world",
			n:      5,
			expect: "hello",
		},
		{
			name:   "multi-byte truncation at codepoint boundary",
			input:  "abc\u00e9def",  // \u00e9 is 2 bytes (é)
			n:      5,              // 'a','b','c' = 3 bytes, 'é' = 2 bytes = 5 bytes total
			expect: "abc\u00e9",
		},
		{
			name:   "multi-byte truncation in middle of codepoint",
			input:  "abc\u00e9def",  // \u00e9 is 2 bytes
			n:      4,              // 3 + 1 byte into é → backs up to 3
			expect: "abc",
		},
		{
			name:   "3-byte utf8 at boundary",
			input:  "ab\u4e16\u754c", // 世界, each 3 bytes; total = 2+3+3 = 8
			n:      5,                // 2 + 3 = 5 → clean cut after 世
			expect: "ab\u4e16",
		},
		{
			name:   "3-byte utf8 mid codepoint",
			input:  "ab\u4e16\u754c",
			n:      6,               // 2 + 3 + 1 byte into 界 → backs up to 5
			expect: "ab\u4e16",
		},
		{
			name:   "4-byte utf8 emoji",
			input:  "hi\U0001F600ok", // 😀 is 4 bytes; total = 2+4+2 = 8
			n:      3,               // 2 + 1 byte into 😀 → backs up to 2
			expect: "hi",
		},
		{
			name:   "4-byte utf8 emoji at boundary",
			input:  "hi\U0001F600ok",
			n:      6,              // 2 + 4 = 6 → clean cut after 😀
			expect: "hi\U0001F600",
		},
		{
			name:   "all multi-byte, n=1 backs up to empty",
			input:  "\u00e9\u00e9", // each 2 bytes
			n:      1,             // 1 byte into first é → backs up to 0
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := truncateUTF8(tt.input, tt.n)
			require.Equal(t, tt.expect, result)
			require.True(t, utf8.ValidString(result), "result should be valid UTF-8")
			require.LessOrEqual(t, len(result), max(tt.n, 0), "result byte length should not exceed n")
		})
	}
}

func TestTruncateMessage(t *testing.T) {
	t.Run("short message unchanged", func(t *testing.T) {
		msg := "short error message"
		require.Equal(t, msg, truncateMessage(msg))
	})

	t.Run("exactly at limit unchanged", func(t *testing.T) {
		msg := strings.Repeat("a", maxMessageLength)
		require.Equal(t, msg, truncateMessage(msg))
	})

	t.Run("over limit is truncated with suffix", func(t *testing.T) {
		msg := strings.Repeat("a", maxMessageLength+100)
		result := truncateMessage(msg)
		require.LessOrEqual(t, len(result), maxMessageLength)
		require.True(t, strings.HasSuffix(result, truncatedMessageSuffix))
	})

	t.Run("one byte over limit is truncated", func(t *testing.T) {
		msg := strings.Repeat("a", maxMessageLength+1)
		result := truncateMessage(msg)
		require.LessOrEqual(t, len(result), maxMessageLength)
		require.True(t, strings.HasSuffix(result, truncatedMessageSuffix))
	})

	t.Run("multi-byte chars truncated cleanly", func(t *testing.T) {
		// Fill with 3-byte UTF-8 characters (世 = 3 bytes each)
		msg := strings.Repeat("\u4e16", maxMessageLength) // way over limit
		result := truncateMessage(msg)
		require.LessOrEqual(t, len(result), maxMessageLength)
		require.True(t, strings.HasSuffix(result, truncatedMessageSuffix))
		require.True(t, utf8.ValidString(result), "result should be valid UTF-8")
	})
}
