package mold

var (
	diveTag            = "dive"
	restrictedTagChars = ".[],|=+()`~!@#$%^&*\\\"/?<>{}"
	restrictedTags     = map[string]struct{}{
		diveTag:   {},
		ignoreTag: {},
	}
)
