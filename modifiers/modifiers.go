package modifiers

import (
	"github.com/go-playground/mold/v4"
)

// New returns a modifier with defaults registered
func New() *mold.Transformer {
	mod := mold.New()
	mod.Register("camel", camelCase)
	mod.Register("default", defaultValue)
	mod.Register("empty", empty)
	mod.Register("lcase", toLower)
	mod.Register("ltrim", trimLeft)
	mod.Register("name", nameCase)
	mod.Register("rtrim", trimRight)
	mod.Register("set", setValue)
	mod.Register("snake", snakeCase)
	mod.Register("slug", slugCase)
	mod.Register("strip_alpha_unicode", stripAlphaUnicodeCase)
	mod.Register("strip_alpha", stripAlphaCase)
	mod.Register("strip_num_unicode", stripNumUnicodeCase)
	mod.Register("strip_num", stripNumCase)
	mod.Register("strip_punctuation", stripPunctuation)
	mod.Register("substr", subStr)
	mod.Register("title", titleCase)
	mod.Register("tprefix", trimPrefix)
	mod.Register("trim", trimSpace)
	mod.Register("tsuffix", trimSuffix)
	mod.Register("ucase", toUpper)
	mod.Register("ucfirst", uppercaseFirstCharacterCase)
	mod.SetTagName("mod")
	return mod
}
