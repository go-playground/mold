package scrubbers

import "github.com/go-playground/mold"

// New returns a scrubber with defaults registered
func New() *mold.Transformer {
	scrub := mold.New()
	scrub.SetTagName("scrub")
	scrub.Register("emails", Emails)
	scrub.Register("name", FullName)
	return scrub
}
