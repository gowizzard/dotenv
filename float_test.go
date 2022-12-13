// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv_test

import (
	"github.com/gowizzard/dotenv/v2"
	"os"
	"reflect"
	"testing"
)

// TestFloat is to test the Float function with table driven tests.
func TestFloat(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    string
		size     int
		set      bool
		expected float64
	}{
		{
			name:     "7.5",
			key:      "TEST1",
			value:    "7.5",
			size:     64,
			set:      true,
			expected: 7.5,
		},
		{
			name:     "175.25",
			key:      "TEST2",
			value:    "175.25",
			size:     32,
			set:      true,
			expected: 175.25,
		},
		{
			name:     "ERROR_64",
			key:      "TEST3",
			value:    "error",
			size:     64,
			set:      true,
			expected: 0,
		},
		{
			name:     "ERROR_32",
			key:      "TEST4",
			value:    "error",
			size:     32,
			set:      true,
			expected: 0,
		},
		{
			name:     "PARSE_ERROR_64",
			key:      "TEST5",
			value:    "",
			size:     64,
			set:      false,
			expected: 0,
		},
		{
			name:     "PARSE_ERROR_32",
			key:      "TEST6",
			value:    "",
			size:     32,
			set:      false,
			expected: 0,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			if value.set {

				t.Setenv(value.key, value.value)
				t.Cleanup(func() {
					os.Clearenv()
				})

			}

			result := dotenv.Float(value.key, value.size)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%f\", got \"%f\"", value.expected, result)
			}

		})

	}

}

// BenchmarkFloat is to test the Float function benchmark timing.
func BenchmarkFloat(b *testing.B) {

	key, value, size := "TEST", "955.5", 64

	b.Setenv(key, value)
	b.Cleanup(func() {
		os.Clearenv()
	})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.Float(key, size)
	}

}
