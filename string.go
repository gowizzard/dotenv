// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv

import (
	"syscall"
)

// String is to get look up the environment and return value as a string if available.
func String(key string) string {

	value, ok := syscall.Getenv(key)
	if ok {
		return value
	}

	return ""

}
