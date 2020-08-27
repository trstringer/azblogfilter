package cache

import (
	"testing"

	"github.com/trstringer/azblogfilter/internal/access"
)

type fakeFileReader struct{}

func (f fakeFileReader) ReadFile(path string) ([]byte, error) {
	return []byte(access.DateTimeLayout), nil
}

func TestLastCachedTime(t *testing.T) {
	lastCachedTime, err := LastCachedTime("/dev/null", fakeFileReader{})
	if err != nil {
		t.Fatal(err)
	}

	if lastCachedTime.IsZero() {
		t.Fatalf("Uxpected zero time for cache")
	}
}
