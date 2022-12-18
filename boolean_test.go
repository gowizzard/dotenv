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

// TestBoolean is to test the Boolean function with table driven tests.
func TestBoolean(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    string
		set      bool
		expected bool
	}{
		{
			name:     "TRUE",
			key:      "TEST1",
			value:    "true",
			set:      true,
			expected: true,
		},
		{
			name:     "FALSE",
			key:      "TEST2",
			value:    "false",
			set:      true,
			expected: false,
		},
		{
			name:     "PARSE_ERROR",
			key:      "TEST3",
			value:    "error",
			set:      true,
			expected: false,
		},
		{
			name:     "NOT_SET",
			key:      "",
			value:    "",
			set:      false,
			expected: false,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			if value.set {
				t.Setenv(value.key, value.value)
			}

			result := dotenv.Boolean(value.key)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%v\", got \"%v\"", value.expected, result)
			}

		})

	}

	t.Cleanup(func() {
		os.Clearenv()
	})

}

// BenchmarkBoolean is to test the Boolean function benchmark timing.
func BenchmarkBoolean(b *testing.B) {

	key, value := "TEST", "true"

	b.Setenv(key, value)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.Boolean(key)
	}

	b.Cleanup(func() {
		os.Clearenv()
	})

}
