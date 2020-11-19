package scrubbers

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func hashString(input string) string {
	h := sha1.New()
	_, _ = io.WriteString(h, input)
	return fmt.Sprintf("%x", h.Sum(nil))
}
