package modifiers

import "github.com/go-playground/mold"

// New returns a modifier with defaults registered
func New() *mold.Transformer {
	scrub := mold.New()
	scrub.SetTagName("mod")
	scrub.Register("trim", TrimSpace)
	scrub.Register("ltrim", TrimLeft)
	scrub.Register("rtrim", TrimRight)
	scrub.Register("tprefix", TrimPrefix)
	scrub.Register("tsuffix", TrimSuffix)
	scrub.Register("lcase", ToLower)
	scrub.Register("ucase", ToUpper)
	scrub.Register("snake", SnakeCase)
	return scrub
}
