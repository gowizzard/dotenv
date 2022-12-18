// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv_test

import (
	"github.com/gowizzard/dotenv/v2"
	"reflect"
	"testing"
)

// TestInteger is to test the Integer function with table driven tests.
func TestInteger(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    string
		base     int
		size     int
		set      bool
		expected int64
	}{
		{
			name:     "0",
			key:      "TEST1",
			value:    "0",
			base:     10,
			size:     8,
			set:      true,
			expected: 0,
		},
		{
			name:     "25",
			key:      "TEST2",
			value:    "25",
			base:     10,
			size:     16,
			set:      true,
			expected: 25,
		},
		{
			name:     "PARSE_ERROR",
			key:      "TEST3",
			base:     10,
			size:     32,
			value:    "error",
			set:      true,
			expected: 0,
		},
		{
			name:     "NOT_SET",
			key:      "TEST4",
			base:     10,
			size:     64,
			value:    "",
			set:      false,
			expected: 0,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			if value.set {
				t.Setenv(value.key, value.value)
			}

			result := dotenv.Integer(value.key, value.base, value.size)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%d\", got \"%d\"", value.expected, result)
			}

		})

	}

}

// BenchmarkInteger is to test the Integer function benchmark timing.
func BenchmarkInteger(b *testing.B) {

	key, value, base, size := "TEST", "175", 10, 64

	b.Setenv(key, value)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.Integer(key, base, size)
	}

}
