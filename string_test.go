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

// TestString is to test the String function with table driven tests.
func TestString(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    string
		expected string
	}{
		{
			name:     "TEST1=\"value\"",
			key:      "TEST1",
			value:    "value",
			expected: "value",
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			err := os.Setenv(value.key, value.value)
			if err != nil {
				t.Error(err)
			}
			defer os.Unsetenv(value.key)

			result := dotenv.String(value.key)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%s\", got \"%s\"", value.expected, result)
			}

		})

	}

}

// BenchmarkString is to test the String function benchmark timing.
func BenchmarkString(b *testing.B) {

	key, value := "TEST", "value"

	err := os.Setenv(key, value)
	if err != nil {
		b.Error(err)
	}
	defer os.Unsetenv(key)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.String(key)
	}

}
