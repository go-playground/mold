package modifiers

import "github.com/go-playground/mold"

// New returns a modifier with defaults registered
func New() *mold.Transformer {
	scrub := mold.New()
	scrub.SetTagName("conform")
	scrub.Register("trimspace", TrimSpace)
	return scrub
}
