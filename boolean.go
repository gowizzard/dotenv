// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package dotenv is used to check the environment
// variables for availability and to return them in different
// data types, so that they can be used directly converted.
package dotenv

import (
	"strconv"
	"syscall"
)

// Boolean is to get look up the environment and return is as boolean.
func Boolean(key string) bool {

	value, ok := syscall.Getenv(key)
	if ok {

		boolean, err := strconv.ParseBool(value)
		if err != nil {
			return false
		}

		return boolean

	}

	return false

}
