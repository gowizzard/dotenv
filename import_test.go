// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv_test

import (
	"github.com/gowizzard/dotenv"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestImport is to test the Import function with table driven tests.
func TestImport(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		perm     os.FileMode
		data     []byte
		expected map[string]string
	}{
		{
			name: "TEST1",
			path: filepath.Join(os.TempDir(), ".env"),
			perm: 0666,
			data: []byte("URL=https://www.test.com/api/v1\nTOKEN=mySecret123456789"),
			expected: map[string]string{
				"URL":   "https://www.test.com/api/v1",
				"TOKEN": "mySecret123456789",
			},
		},
		{
			name: "TEST2",
			path: filepath.Join(os.TempDir(), ".env"),
			perm: 0666,
			data: []byte("SMTP_HOST='mail.mailcow.org'\nSMTP_PORT='587'\nSMTP_USERNAME='hello@mailcow.org'\nSMTP_PASSWORD='test12345'"),
			expected: map[string]string{
				"SMTP_HOST":     "mail.mailcow.org",
				"SMTP_PORT":     "587",
				"SMTP_USERNAME": "hello@mailcow.org",
				"SMTP_PASSWORD": "test12345",
			},
		},
		{
			name: "TEST3",
			path: filepath.Join(os.TempDir(), ".env"),
			perm: 0666,
			data: []byte("# Language settings\nLANGUAGE=\"de_DE\""),
			expected: map[string]string{
				"LANGUAGE": "de_DE",
			},
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			err := os.WriteFile(value.path, value.data, value.perm)
			if err != nil {
				t.Error(err)
			}

			err = dotenv.Import(value.path)
			if err != nil {
				t.Error(err)
			}

			for index, value := range value.expected {

				result := os.Getenv(index)

				if !reflect.DeepEqual(value, result) {
					t.Errorf("expected: \"%s\", got \"%s\"", value, result)
				}

			}

		})

	}

}

// BenchmarkImport is to test the Import function benchmark timing.
func BenchmarkImport(b *testing.B) {

	path, perm, data := filepath.Join(os.TempDir(), ".env"), os.FileMode(0666), []byte("USERNAME=\"gowizzard\"\nREPO=\"dotenv\"")

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

}
