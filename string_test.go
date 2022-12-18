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

// TestString is to test the String function with table driven tests.
func TestString(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    string
		set      bool
		expected string
	}{
		{
			name:     "VALUE",
			key:      "TEST1",
			value:    "value",
			set:      true,
			expected: "value",
		},
		{
			name:     "NOT_SET",
			key:      "TEST2",
			value:    "",
			set:      false,
			expected: "",
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			if value.set {
				t.Setenv(value.key, value.value)
			}

			result := dotenv.String(value.key)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%s\", got \"%s\"", value.expected, result)
			}

		})

	}

	t.Cleanup(func() {
		os.Clearenv()
	})

}

// BenchmarkString is to test the String function benchmark timing.
func BenchmarkString(b *testing.B) {

	key, value := "TEST", "value"

	b.Setenv(key, value)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dotenv.String(key)
	}

	b.Cleanup(func() {
		os.Clearenv()
	})

}
