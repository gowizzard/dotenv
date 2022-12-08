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

// TestInteger is to test the Integer function with table driven tests.
func TestInteger(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    string
		expected int
	}{
		{
			name:     "TEST1=\"0\"",
			key:      "TEST1",
			value:    "0",
			expected: 0,
		},
		{
			name:     "TEST2=\"25\"",
			key:      "TEST2",
			value:    "25",
			expected: 25,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			err := os.Setenv(value.key, value.value)
			if err != nil {
				t.Error(err)
			}
			defer os.Unsetenv(value.key)

			result := dotenv.Integer(value.key)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%d\", got \"%d\"", value.expected, result)
			}

		})

	}

}

// BenchmarkInteger is to test the Integer function benchmark timing.
func BenchmarkInteger(b *testing.B) {

	key, value := "TEST", "175"

	err := os.Setenv(key, value)
	if err != nil {
		b.Error(err)
	}
	defer os.Unsetenv(key)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.Integer(key)
	}

}
