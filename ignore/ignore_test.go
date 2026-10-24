package ignore

import (
	"strings"
	"testing"
)

func TestIgnoreMatcher(t *testing.T) {
	ignoreFile := `
# ignore this comment
  
dir1/
/root/
prefix
*.go
**/data.*
`
	ig, err := ParseIgnoreFile(strings.NewReader(ignoreFile))
	if err != nil {
		t.Errorf("error in ignoreFile")
	}
	tests := []struct {
		path      string
		wantMatch bool
	}{
		{
			path:      "dir1/readme.md",
			wantMatch: true,
		},
		{
			path:      "dir1/dir2/readme.md",
			wantMatch: true,
		},

		{
			path:      "foo.go",
			wantMatch: true,
		},
		{
			path:      "dir2/foo.go",
			wantMatch: false,
		},
		{
			path:      "dir3/data.xyz",
			wantMatch: true,
		},
		{
			path:      "root/readme.md",
			wantMatch: true,
		},
		{
			path:      "nested/root/readme.md",
			wantMatch: false,
		},
		{
			path:      "prefix-file.txt",
			wantMatch: true,
		},
		{
			path:      "nested/prefix-file.txt",
			wantMatch: false,
		},
		{
			path:      "readme.md",
			wantMatch: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := ig.Match(tt.path); got != tt.wantMatch {
				t.Errorf("got %t, expected %t", got, tt.wantMatch)
			}
		})
	}
}
