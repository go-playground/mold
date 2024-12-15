package mold

import (
	"math"
	"reflect"
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestGetPrimitiveValue(t *testing.T) {
	tests := []struct {
		name        string
		typ         reflect.Kind
		value       string
		expected    reflect.Value
		expectError string
	}{
		{
			name:     "string",
			typ:      reflect.String,
			value:    "test",
			expected: reflect.ValueOf("test"),
		},
		{
			name:     "int",
			typ:      reflect.Int,
			value:    "123",
			expected: reflect.ValueOf(123),
		},
		{
			name:     "int8",
			typ:      reflect.Int8,
			value:    "123",
			expected: reflect.ValueOf(int8(123)),
		},
		{
			name:     "int16",
			typ:      reflect.Int16,
			value:    "123",
			expected: reflect.ValueOf(int16(123)),
		},
		{
			name:     "bool",
			typ:      reflect.Bool,
			value:    "true",
			expected: reflect.ValueOf(true),
		},
		{
			name:        "error while parsing int",
			typ:         reflect.Int,
			value:       "abc",
			expectError: "mold: failed to parse value for type int: strconv.Atoi: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing int8",
			typ:         reflect.Int8,
			value:       "abc",
			expectError: "mold: failed to parse value for type int8: strconv.ParseInt: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing int16",
			typ:         reflect.Int16,
			value:       "abc",
			expectError: "mold: failed to parse value for type int16: strconv.ParseInt: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing int32",
			typ:         reflect.Int32,
			value:       "abc",
			expectError: "mold: failed to parse value for type int32: strconv.ParseInt: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing int64",
			typ:         reflect.Int64,
			value:       "abc",
			expectError: "mold: failed to parse value for type int64: strconv.ParseInt: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing uint",
			typ:         reflect.Uint,
			value:       "abc",
			expectError: "mold: failed to parse value for type uint: strconv.ParseUint: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing uint8",
			typ:         reflect.Uint8,
			value:       "abc",
			expectError: "mold: failed to parse value for type uint8: strconv.ParseUint: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing uint16",
			typ:         reflect.Uint16,
			value:       "12.34",
			expectError: "mold: failed to parse value for type uint16: strconv.ParseUint: parsing \"12.34\": invalid syntax",
		},
		{
			name:        "error while parsing uint32",
			typ:         reflect.Uint32,
			value:       "12.34",
			expectError: "mold: failed to parse value for type uint32: strconv.ParseUint: parsing \"12.34\": invalid syntax",
		},
		{
			name:        "error while parsing uint64",
			typ:         reflect.Uint64,
			value:       "12.34",
			expectError: "mold: failed to parse value for type uint64: strconv.ParseUint: parsing \"12.34\": invalid syntax",
		},
		{
			name:        "error while parsing float32",
			typ:         reflect.Float32,
			value:       "abc",
			expectError: "mold: failed to parse value for type float32: strconv.ParseFloat: parsing \"abc\": invalid syntax",
		},
		{
			name:        "error while parsing bool",
			typ:         reflect.Bool,
			value:       "invalid bool",
			expectError: "mold: failed to parse value for type bool: strconv.ParseBool: parsing \"invalid bool\": invalid syntax",
		},
		{
			name:        "unsupported type",
			typ:         reflect.Struct,
			value:       "abc",
			expectError: "mold: unsupported field type: struct",
		},
		{
			name:     "uint",
			typ:      reflect.Uint,
			value:    "123",
			expected: reflect.ValueOf(uint(123)),
		},
		{
			name:     "uint8",
			typ:      reflect.Uint8,
			value:    "123",
			expected: reflect.ValueOf(uint8(123)),
		},
		{
			name:     "uint16",
			typ:      reflect.Uint16,
			value:    "123",
			expected: reflect.ValueOf(uint16(123)),
		},
		{
			name:     "uint32",
			typ:      reflect.Uint32,
			value:    "123",
			expected: reflect.ValueOf(uint32(123)),
		},
		{
			name:     "uint64",
			typ:      reflect.Uint64,
			value:    "123",
			expected: reflect.ValueOf(uint64(123)),
		},
		{
			name:     "float32",
			typ:      reflect.Float32,
			value:    "123.45",
			expected: reflect.ValueOf(float32(123.45)),
		},
		{
			name:     "float64",
			typ:      reflect.Float64,
			value:    "123.45",
			expected: reflect.ValueOf(float64(123.45)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetPrimitiveValue(tc.typ, tc.value)
			if tc.expectError != "" {
				NotEqual(t, nil, err)
				Equal(t, tc.expectError, err.Error())
			} else {
				Equal(t, nil, err)
				switch tc.typ {
				case reflect.String:
					Equal(t, tc.expected.String(), actual.String())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					Equal(t, tc.expected.Int(), actual.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					Equal(t, tc.expected.Uint(), actual.Uint())
				case reflect.Float32, reflect.Float64:
					// could not assert equal float because of precision issues
					// so we just check in 4 decimal places
					decimalMask := math.Pow(10, 4)
					Equal(t, math.Round(tc.expected.Float()*decimalMask), math.Round(actual.Float()*decimalMask))
				case reflect.Bool:
					Equal(t, tc.expected.Bool(), actual.Bool())
				default:
					Equal(t, tc.expected.Interface(), actual.Interface())
				}
			}
		})
	}
}
