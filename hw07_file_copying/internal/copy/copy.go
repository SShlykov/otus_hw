package copy

import (
	"errors"
	"github.com/SShlykov/otus_hw/hw07_file_copying/internal/progress"
	"io"
	"math"
	"os"
	"path/filepath"
)

var (
	ErrorFileNotFound        = errors.New("file not found")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if err := validateParams(fromPath, toPath, offset, limit); err != nil {
		return err
	}

	source, err := os.OpenFile(fromPath, os.O_RDONLY, 0444)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrorFileNotFound
		}
		return err
	}
	defer source.Close()

	sourceInfo, err := source.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	sourceLen := sourceInfo.Size()
	bytesToCopy := sourceLen - offset
	if limit != 0 && limit < bytesToCopy {
		bytesToCopy = limit
	}

	if offset > sourceLen {
		return ErrOffsetExceedsFileSize
	}

	var target *os.File
	target, err = os.Create(toPath)
	if err != nil {
		return err
	}
	defer target.Close()

	return doCopy(source, target, offset, bytesToCopy)
}

func doCopy(source *os.File, target *os.File, offset, bytesToCopy int64) error {
	bar := progress.NewBar("copy file", bytesToCopy)
	defer bar.Finish()

	bufferSize := getBufferSize(bytesToCopy)
	buf := make([]byte, bufferSize)
	current := int64(0)
	if _, err := source.Seek(offset, 0); err != nil {
		return err
	}

	for current < bytesToCopy {
		if _, err := source.Read(buf); err != nil && err != io.EOF {
			return err
		}

		remained := min(bytesToCopy-current, bufferSize)
		str := buf[:remained]

		if _, err := target.Write(str); err != nil {
			return err
		}
		current += remained
		bar.Add(remained)
	}

	return nil
}

func getBufferSize(bytesToCopy int64) int64 {
	if bytesToCopy < 100 {
		return 4
	} else {
		return int64(math.Pow(2, math.Ceil(math.Log(float64(bytesToCopy/100))/math.Log(2))))
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

var (
	ErrorFromPathEmpty = errors.New("from path is empty")
	ErrorToPathEmpty   = errors.New("to path is empty")
	ErrorSamePaths     = errors.New("from and to paths are the same")
	ErrorNegativeValue = errors.New("offset and limit should be positive")
)

func validateParams(fromPath, toPath string, offset, limit int64) error {
	if fromPath == "" {
		return ErrorFromPathEmpty
	}
	if toPath == "" {
		return ErrorToPathEmpty
	}
	cleanFromPath := filepath.Clean(fromPath)
	cleanToPath := filepath.Clean(toPath)

	if cleanFromPath == cleanToPath {
		return ErrorSamePaths
	}
	if offset < 0 || limit < 0 {
		return ErrorNegativeValue
	}

	return nil
}
