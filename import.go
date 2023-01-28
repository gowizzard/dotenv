// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv

import (
	"syscall"
)

// Import is read the environment variable file and use regex to find
// all sub matches. After that we initialize the environment variables to local.
func Import(path string) error {

	matches, err := read(path)
	if err != nil {
		return err
	}

	for index, value := range matches {
		err = syscall.Setenv(index, value)
		if err != nil {
			return err
		}
	}

	return nil

}
