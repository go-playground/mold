package modifiers

import (
	"context"
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
)

func TestDefaultSetSpecialTypes(t *testing.T) {
	conform := New()

	tests := []struct {
		name        string
		field       interface{}
		tags        string
		vf          func(field interface{})
		expectError bool
	}{
		{
			name:  "default map",
			field: (map[string]struct{})(nil),
			tags:  "default",
			vf: func(field interface{}) {
				m := field.(map[string]struct{})
				Equal(t, len(m), 0)
			},
		},
		{
			name:  "default map with size",
			field: (map[string]struct{})(nil),
			tags:  "default=5",
			vf: func(field interface{}) {
				m := field.(map[string]struct{})
				Equal(t, len(m), 0)
			},
		},
		{
			name:  "set map with size",
			field: (map[string]struct{})(nil),
			tags:  "set=5",
			vf: func(field interface{}) {
				m := field.(map[string]struct{})
				Equal(t, len(m), 0)
			},
		},
		{
			name:  "default slice",
			field: ([]string)(nil),
			tags:  "default",
			vf: func(field interface{}) {
				m := field.([]string)
				Equal(t, len(m), 0)
				Equal(t, cap(m), 0)
			},
		},
		{
			name:  "default slice with capacity",
			field: ([]string)(nil),
			tags:  "default=5",
			vf: func(field interface{}) {
				m := field.([]string)
				Equal(t, len(m), 0)
				Equal(t, cap(m), 5)
			},
		},
		{
			name:  "set slice",
			field: ([]string)(nil),
			tags:  "set",
			vf: func(field interface{}) {
				m := field.([]string)
				Equal(t, len(m), 0)
				Equal(t, cap(m), 0)
			},
		},
		{
			name:  "set slice with capacity",
			field: ([]string)(nil),
			tags:  "set=5",
			vf: func(field interface{}) {
				m := field.([]string)
				Equal(t, len(m), 0)
				Equal(t, cap(m), 5)
			},
		},
		{
			name:  "default chan",
			field: (chan struct{})(nil),
			tags:  "default",
			vf: func(field interface{}) {
				m := field.(chan struct{})
				Equal(t, len(m), 0)
				Equal(t, cap(m), 0)
			},
		},
		{
			name:  "default chan with buffer",
			field: (chan struct{})(nil),
			tags:  "default=5",
			vf: func(field interface{}) {
				m := field.(chan struct{})
				Equal(t, len(m), 0)
				Equal(t, cap(m), 5)
			},
		},
		{
			name:  "default time.Time",
			field: time.Time{},
			tags:  "default",
			vf: func(field interface{}) {
				m := field.(time.Time)
				Equal(t, m.Location(), time.Local)
			},
		},
		{
			name:  "default time.Time utc",
			field: time.Time{},
			tags:  "default=utc",
			vf: func(field interface{}) {
				m := field.(time.Time)
				Equal(t, m.Location(), time.UTC)
			},
		},
		{
			name:  "default time.Time to value",
			field: time.Time{},
			tags:  "default=2023-05-28T15:50:31Z",
			vf: func(field interface{}) {
				m := field.(time.Time)
				Equal(t, m.Location(), time.UTC)

				tm, err := time.Parse(time.RFC3339Nano, "2023-05-28T15:50:31Z")
				Equal(t, err, nil)
				Equal(t, tm.Equal(m), true)

			},
		},
		{
			name:  "set time.Time",
			field: time.Time{},
			tags:  "set",
			vf: func(field interface{}) {
				m := field.(time.Time)
				Equal(t, m.Location(), time.Local)
			},
		},
		{
			name:  "set time.Time utc",
			field: time.Time{},
			tags:  "set=utc",
			vf: func(field interface{}) {
				m := field.(time.Time)
				Equal(t, m.Location(), time.UTC)
			},
		},
		{
			name:  "set time.Time to value",
			field: time.Time{},
			tags:  "set=2023-05-28T15:50:31Z",
			vf: func(field interface{}) {
				m := field.(time.Time)
				Equal(t, m.Location(), time.UTC)

				tm, err := time.Parse(time.RFC3339Nano, "2023-05-28T15:50:31Z")
				Equal(t, err, nil)
				Equal(t, tm.Equal(m), true)

			},
		},
		{
			name:  "default pointer to slice",
			field: (*[]string)(nil),
			tags:  "default",
			vf: func(field interface{}) {
				m := field.([]string)
				Equal(t, len(m), 0)
			},
		},
		{
			name:  "set pointer to slice",
			field: (*[]string)(nil),
			tags:  "set",
			vf: func(field interface{}) {
				m := field.([]string)
				Equal(t, len(m), 0)
			},
		}, {
			name:  "default pointer to int",
			field: (*int)(nil),
			tags:  "default=5",
			vf: func(field interface{}) {
				m := field.(int)
				Equal(t, m, 5)
			},
		},
		{
			name:  "default pointer to string",
			field: (*string)(nil),
			tags:  "default=test",
			vf: func(field interface{}) {
				m := field.(string)
				Equal(t, m, "test")
			},
		},
		{
			name:  "set pointer to string",
			field: (*string)(nil),
			tags:  "set",
			vf: func(field interface{}) {
				m := field.(string)
				Equal(t, m, "")
			},
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
			tc.vf(tc.field)
		})
	}
}

func TestSet(t *testing.T) {

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
			name:     "set State (although enum default value should be the default in practice)",
			field:    state,
			tags:     "set=5",
			expected: FINISHED,
		},
		{
			name:     "set string",
			field:    "",
			tags:     "set=test",
			expected: "test",
		},
		{
			name:     "set string",
			field:    "existing_value",
			tags:     "set=test",
			expected: "test",
		},
		{
			name:     "set int",
			field:    0,
			tags:     "set=3",
			expected: 3,
		},
		{
			name:     "set uint",
			field:    uint(0),
			tags:     "default=4",
			expected: uint(4),
		},
		{
			name:     "set float",
			field:    float32(0),
			tags:     "set=5",
			expected: float32(5),
		},
		{
			name:     "set bool",
			field:    false,
			tags:     "set=true",
			expected: true,
		},
		{
			name:     "set time.Duration",
			field:    time.Duration(0),
			tags:     "set=1s",
			expected: time.Duration(1_000_000_000),
		},
		{
			name:        "bad set time.Duration",
			field:       time.Duration(0),
			tags:        "set=rex",
			expectError: true,
		},
		{
			name:        "set default int",
			field:       0,
			tags:        "set=abc",
			expectError: true,
		},
		{
			name:        "bad set uint",
			field:       uint(0),
			tags:        "set=abc",
			expectError: true,
		},
		{
			name:        "bad set float",
			field:       float32(0),
			tags:        "default=abc",
			expectError: true,
		},
		{
			name:        "bad set bool",
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
			name:     "default string",
			field:    "existing_value",
			tags:     "default=test",
			expected: "existing_value",
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
		{
			name:     "default nil pointer to int",
			field:    (*int)(nil),
			tags:     "default=3",
			expected: 3,
		},
		{
			name:     "default not nil pointer to int",
			field:    newPointer(1),
			tags:     "default=3",
			expected: 1,
		},
		{
			name:     "default nil pointer to string",
			field:    (*string)(nil),
			tags:     "default=test",
			expected: "test",
		},
		{
			name:     "default not nil pointer to string",
			field:    newPointer("existing_value"),
			tags:     "default=test",
			expected: "existing_value",
		},
		{
			name:     "default nil pointer to bool",
			field:    (*bool)(nil),
			tags:     "default=true",
			expected: true,
		},
		{
			name:     "default not nil pointer to bool",
			field:    newPointer(true),
			tags:     "default=true",
			expected: true,
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

func TestEmpty(t *testing.T) {

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
			name:     "empty enum",
			field:    FINISHED,
			tags:     "empty",
			expected: state,
		},
		{
			name:     "empty string",
			field:    "test",
			tags:     "empty",
			expected: "",
		},
		{
			name:     "empty int",
			field:    10,
			tags:     "empty",
			expected: 0,
		},
		{
			name:     "empty uint",
			field:    uint(10),
			tags:     "empty",
			expected: uint(0),
		},
		{
			name:     "empty float",
			field:    float32(10),
			tags:     "empty",
			expected: float32(0),
		},
		{
			name:     "empty bool",
			field:    true,
			tags:     "empty",
			expected: false,
		},
		{
			name:     "empty time.Duration",
			field:    time.Duration(10),
			tags:     "empty",
			expected: time.Duration(0),
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

func newPointer[T any](value T) *T {
	return &value
}
