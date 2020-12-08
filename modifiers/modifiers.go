package modifiers

import (
	"github.com/go-playground/mold/v3"
)

// New returns a modifier with defaults registered
func New() *mold.Transformer {
	mod := mold.New()
	mod.SetTagName("mod")
	mod.Register("trim", TrimSpace)
	mod.Register("ltrim", TrimLeft)
	mod.Register("rtrim", TrimRight)
	mod.Register("tprefix", TrimPrefix)
	mod.Register("tsuffix", TrimSuffix)
	mod.Register("lcase", ToLower)
	mod.Register("ucase", ToUpper)
	mod.Register("snake", SnakeCase)
	mod.Register("alpha", AlphaCase)
	mod.Register("notalpha", NotAlphaCase)
	mod.Register("default", Default)
	return mod
}
