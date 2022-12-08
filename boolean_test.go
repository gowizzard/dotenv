// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv_test

import (
	"github.com/gowizzard/dotenv"
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
		expected bool
	}{
		{
			name:     "PRODUCTION=\"true\"",
			key:      "PRODUCTION",
			value:    "true",
			expected: true,
		},
		{
			name:     "DEVELOPMENT=\"false\"",
			key:      "DEVELOPMENT",
			value:    "false",
			expected: false,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			err := os.Setenv(value.key, value.value)
			if err != nil {
				t.Error(err)
			}
			defer os.Unsetenv(value.key)

			result := dotenv.Boolean(value.key)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%v\", got \"%v\"", value.expected, result)
			}

		})

	}

}

// BenchmarkBoolean is to test the Boolean function benchmark timing.
func BenchmarkBoolean(b *testing.B) {

	key, value := "DEVELOPMENT", "true"

	err := os.Setenv(key, value)
	if err != nil {
		b.Error(err)
	}
	defer os.Unsetenv(key)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.Boolean(key)
	}

}
