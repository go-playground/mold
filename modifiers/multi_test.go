package modifiers

import (
	"context"
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
)

func TestDefault(t *testing.T) {

	type State int
	const FINISHED State = 5

	var state State

	conform := New()

	tests := []struct {
		name        string
		field       interface{}
		tags        string
		expected    interface{}
		expectError bool
	}{
		{
			name:     "default State (although enum default value should be the default in practice)",
			field:    state,
			tags:     "default=5",
			expected: FINISHED,
		},
		{
			name:     "default string",
			field:    "",
			tags:     "default=test",
			expected: "test",
		},
		{
			name:     "default int",
			field:    0,
			tags:     "default=3",
			expected: 3,
		},
		{
			name:     "default uint",
			field:    uint(0),
			tags:     "default=4",
			expected: uint(4),
		},
		{
			name:     "default float",
			field:    float32(0),
			tags:     "default=5",
			expected: float32(5),
		},
		{
			name:     "default bool",
			field:    false,
			tags:     "default=true",
			expected: true,
		},
		{
			name:     "default time.Duration",
			field:    time.Duration(0),
			tags:     "default=1s",
			expected: time.Duration(1_000_000_000),
		},
		{
			name:        "bad default time.Duration",
			field:       time.Duration(0),
			tags:        "default=rex",
			expectError: true,
		},
		{
			name:        "bad default int",
			field:       0,
			tags:        "default=abc",
			expectError: true,
		},
		{
			name:        "bad default uint",
			field:       uint(0),
			tags:        "default=abc",
			expectError: true,
		},
		{
			name:        "bad default float",
			field:       float32(0),
			tags:        "default=abc",
			expectError: true,
		},
		{
			name:        "bad default bool",
			field:       false,
			tags:        "default=blue",
			expectError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := conform.Field(context.Background(), &tc.field, tc.tags)
			if tc.expectError {
				NotEqual(t, err, nil)
				return
			}
			Equal(t, err, nil)
			Equal(t, tc.field, tc.expected)
		})
	}
}
