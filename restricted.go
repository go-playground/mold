package mold

const (
	diveTag            = "dive"
	restrictedTagChars = ".[],|=+()`~!@#$%^&*\\\"/?<>{}"
)

var (
	restrictedTags = map[string]struct{}{
		diveTag:   {},
		ignoreTag: {},
	}
)
