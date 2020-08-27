package cache

import (
	"io/ioutil"
	"time"

	"github.com/trstringer/azblogfilter/internal/access"
)

// FileReader is the interface to be able to read a file.
type FileReader interface {
	ReadFile(string) ([]byte, error)
}

// FileSystemReader is the FileReader for the filesystem.
type FileSystemReader struct{}

// ReadFile reads a file from the filesystem.
func (f FileSystemReader) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// LastCachedTime returns the time that the cache was last updated.
func LastCachedTime(path string, fileReader FileReader) (time.Time, error) {
	contentBytes, err := fileReader.ReadFile(path)
	if err != nil {
		return time.Time{}, err
	}

	contents := string(contentBytes)
	lastCachedTime, err := time.Parse(access.DateTimeLayout, contents)
	if err != nil {
		return time.Time{}, err
	}

	return lastCachedTime, nil
}

// LastCachedTimeFromFileSystem returns the last cached time from disk.
func LastCachedTimeFromFileSystem(path string) (time.Time, error) {
	return LastCachedTime(path, FileSystemReader{})
}

// SetLastCachedTime sets the cache for the last time.
func SetLastCachedTime(path string, lastUpdated time.Time) error {
	lastUpdatedFormatted := lastUpdated.Format(access.DateTimeLayout)
	err := ioutil.WriteFile(path, []byte(lastUpdatedFormatted), 755)
	if err != nil {
		return err
	}
	return nil
}
