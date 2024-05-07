package copy

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	const inputFile = "test_input/input.txt"
	const outputFile = "test_input/output.txt"

	tests := []struct {
		name    string
		input   string
		output  string
		offset  int64
		limit   int64
		wantErr bool
		err     error
	}{
		{"Copy without offset and limit", inputFile, outputFile, 0, 0, false, nil},
		{"Copy with offset", inputFile, outputFile, 100, 0, false, nil},
		{"Copy with limit", inputFile, outputFile, 0, 100, false, nil},
		{"Copy with offset and limit", inputFile, outputFile, 100, 100, false, nil},
		{"With Error", inputFile, inputFile, 100, 100, false, ErrorSamePaths},
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
