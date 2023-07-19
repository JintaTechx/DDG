package writer

import (
	"io"
	"os"
)

// CreateFile creates a new file based in the path passed by parameter
func CreateFile(path string) (io.Writer, error) {
	file, err := os.Create(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}
