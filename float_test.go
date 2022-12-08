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

// TestFloat is to test the Float function with table driven tests.
func TestFloat(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    string
		bit      int
		expected float64
	}{
		{
			name:     "TEST1=\"7.5\"",
			key:      "TEST1",
			value:    "7.5",
			bit:      64,
			expected: 7.5,
		},
		{
			name:     "TEST2=\"175.25\"",
			key:      "TEST2",
			value:    "175.25",
			bit:      32,
			expected: 175.25,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			err := os.Setenv(value.key, value.value)
			if err != nil {
				t.Error(err)
			}
			defer os.Unsetenv(value.key)

			result := dotenv.Float(value.key, value.bit)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%f\", got \"%f\"", value.expected, result)
			}

		})

	}

}

// BenchmarkFloat is to test the Float function benchmark timing.
func BenchmarkFloat(b *testing.B) {

	key, value, bit := "TEST", "955.5", 64

	err := os.Setenv(key, value)
	if err != nil {
		b.Error(err)
	}
	defer os.Unsetenv(key)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.Float(key, bit)
	}

}
