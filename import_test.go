// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv_test

import (
	"github.com/gowizzard/dotenv/v2"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
	"testing"
)

// TestImport is to test the Import function with table driven tests.
func TestImport(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		perm     os.FileMode
		data     []byte
		write    bool
		error    bool
		expected map[string]string
	}{
		{
			name:  "WITHOUT_QUOTES",
			path:  filepath.Join(t.TempDir(), ".env"),
			perm:  os.ModePerm,
			data:  []byte("TEST1=value\nTEST2=25"),
			write: true,
			error: false,
			expected: map[string]string{
				"TEST1": "value",
				"TEST2": "25",
			},
		},
		{
			name:  "SINGLE_QUOTES",
			path:  filepath.Join(t.TempDir(), ".env"),
			perm:  os.ModePerm,
			data:  []byte("TEST1='value'\nTEST2='25'\nTEST3='42.5'\nTEST4='true'"),
			write: true,
			error: false,
			expected: map[string]string{
				"TEST1": "value",
				"TEST2": "25",
				"TEST3": "42.5",
				"TEST4": "true",
			},
		},
		{
			name:  "DOUBLE_QUOTES",
			path:  filepath.Join(t.TempDir(), ".env"),
			perm:  os.ModePerm,
			data:  []byte("# This is a test command.\nTEST1=\"value\""),
			write: true,
			error: false,
			expected: map[string]string{
				"TEST1": "value",
			},
		},
		{
			name:     "NO_MATCHES",
			path:     filepath.Join(t.TempDir(), ".env"),
			perm:     os.ModePerm,
			data:     []byte("=value\n=25"),
			write:    true,
			error:    true,
			expected: nil,
		},
		{
			name:     "FILE_NOT_EXIST",
			path:     "",
			perm:     os.ModePerm,
			data:     nil,
			write:    false,
			error:    true,
			expected: nil,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			if value.write {

				err := os.WriteFile(value.path, value.data, value.perm)
				if err != nil {
					t.Error(err)
				}

			}

			err := dotenv.Import(value.path)
			if err != nil && !value.error {
				t.Error(err)
			}

			if value.error {
				return
			}

			for index, value := range value.expected {

				result, ok := syscall.Getenv(index)
				if ok && !reflect.DeepEqual(value, result) {
					t.Errorf("expected: \"%s\", got \"%s\"", value, result)
				}

			}

		})

	}

	t.Cleanup(func() {
		syscall.Clearenv()
	})

}

// BenchmarkImport is to test the Import function benchmark timing.
func BenchmarkImport(b *testing.B) {

	path, perm, data := filepath.Join(b.TempDir(), ".env"), os.ModePerm, []byte("TEST1=\"value\"\nTEST2=\"value\"")

	err := os.WriteFile(path, data, perm)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := dotenv.Import(path)
		if err != nil {
			b.Error(err)
		}
	}

	b.Cleanup(func() {
		syscall.Clearenv()
	})

}
