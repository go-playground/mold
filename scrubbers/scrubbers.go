package scrubbers

import "gopkg.in/go-playground/mold.v2"

// New returns a scrubber with defaults registered
func New() *mold.Transformer {
	scrub := mold.New()
	scrub.SetTagName("scrub")
	scrub.Register("emails", Emails)
	scrub.Register("text", textFn("text"))
	scrub.Register("email", textFn("email"))
	scrub.Register("name", textFn("name"))
	scrub.Register("fname", textFn("fname"))
	scrub.Register("lname", textFn("lname"))
	return scrub
}
