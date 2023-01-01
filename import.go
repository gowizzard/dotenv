// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv

import (
	"os"
	"regexp"
	"syscall"
)

// data is to save the environ data.
type data struct {
	Key   []byte
	Value []byte
}

// regex is to save the compiled expression.
var regex = regexp.MustCompile(`(?m)^(?P<key>\w+?)=["']?(?P<value>.*)(?:["']|\b)$`)

// Import is read the environment variable file and use regex to find
// all sub matches. After that we initialize the environment variables to local.
func Import(path string) error {

	read, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	for _, value := range regex.FindAllSubmatch(read, -1) {

		environ := new(data)

		names := regex.SubexpNames()
		for index := range names {
			switch names[index] {
			case "key":
				environ.Key = value[index]
			case "value":
				environ.Value = value[index]
			}
		}

		err = syscall.Setenv(string(environ.Key), string(environ.Value))
		if err != nil {
			return err
		}

	}

	return nil

}
