// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv

import (
	"errors"
	"os"
	"regexp"
	"syscall"
)

// data is to save the key and value of the environment variable.
type data struct {
	key   []byte
	value []byte
}

// regex is to save the compiled expression.
var regex = regexp.MustCompile(`(?m)^(?P<key>\w+?)=(?:["']|\b)(?P<value>.+?)(?:["']|\b)$`)

// Import is read the environment variable file and use regex to find
// all sub matches. After that we initialize the environment variables to local.
func Import(path string) error {

	read, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if len(read) == 0 {
		return errors.New("file is empty")
	}

	matches := regex.FindAllSubmatch(read, -1)
	if matches == nil {
		return errors.New("no matches found")
	}

	for _, value := range matches {

		variable := new(data)

		names := regex.SubexpNames()
		for index := range names {
			switch names[index] {
			case "key":
				variable.key = value[index]
			case "value":
				variable.value = value[index]
			}
		}

		err = syscall.Setenv(string(variable.key), string(variable.value))
		if err != nil {
			return err
		}

	}

	return nil

}
