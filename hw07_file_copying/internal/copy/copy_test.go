package copy

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	const inputFile = "test_input/input.txt"
	const outputFile = "tmp/output.txt"

	tests := []struct {
		name    string
		offset  int64
		limit   int64
		wantErr bool
		err     error
	}{
		{"Copy without offset and limit", 0, 0, false, nil},
		{"Copy with offset", 100, 0, false, nil},
		{"Copy with limit", 0, 100, false, nil},
		{"Copy with offset and limit", 100, 100, false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(inputFile, outputFile, tt.offset, tt.limit)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
			_ = os.Remove(outputFile)
		})
	}
}
