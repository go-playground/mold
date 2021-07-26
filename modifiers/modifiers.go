package modifiers

import (
	"github.com/go-playground/mold/v4"
)

// New returns a modifier with defaults registered
func New() *mold.Transformer {
	mod := mold.New()
	mod.SetTagName("mod")
	mod.Register("trim", trimSpace)
	mod.Register("ltrim", trimLeft)
	mod.Register("rtrim", trimRight)
	mod.Register("tprefix", trimPrefix)
	mod.Register("tsuffix", trimSuffix)
	mod.Register("lcase", toLower)
	mod.Register("ucase", toUpper)
	mod.Register("snake", snakeCase)
	mod.Register("title", titleCase)
	mod.Register("name", nameCase)
	mod.Register("ucfirst", uppercaseFirstCharacterCase)
	mod.Register("strip_alpha", stripAlphaCase)
	mod.Register("strip_num", stripNumCase)
	mod.Register("strip_num_unicode", stripNumUnicodeCase)
	mod.Register("strip_alpha_unicode", stripAlphaUnicodeCase)
	mod.Register("strip_punctuation", stripPunctuation)
	mod.Register("camel", camelCase)
	mod.Register("default", defaultValue)
	return mod
}
