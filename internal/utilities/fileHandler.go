/*
Ensures that font file is ready for use and load the file else it downloads this file
*/
package ascii

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Make sure the file exist if not it downloads the file to the project
func EnsureFile(filename, url string) error {
	if FileExists(filename) {
		return nil
	}

	data, err := DownloadFile(url)
	if err != nil {
		return err
	}

	return SaveFile(filename, data)
}
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// downloads ascii font file
func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func SaveFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0o644)
}

// Read and split the file contents
func LoadAsciiFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
