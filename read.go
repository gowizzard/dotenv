// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv

import (
	"errors"
	"os"
	"regexp"
)

// regex is to save the compiled expression.
var regex = regexp.MustCompile(`(?m)^(?P<key>\w+?)=(?:["']|\b)(?P<value>.+?)(?:["']|\b)$`)

// read is to read the environment variable file and use regex to find all sub matches.
func read(path string) (map[string]string, error) {

	read, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(read) == 0 {
		return nil, errors.New("file is empty")
	}

	matches := regex.FindAllSubmatch(read, -1)
	if matches == nil {
		return nil, errors.New("no matches found")
	}

	var environ = make(map[string]string)
	for _, value := range matches {

		index := map[string]int{
			"key":   regex.SubexpIndex("key"),
			"value": regex.SubexpIndex("value"),
		}

		environ[string(value[index["key"]])] = string(value[index["value"]])

	}

	return environ, nil

}
