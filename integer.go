// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv

import (
	"os"
	"strconv"
)

// Integer is to get look up the environment and return is as int64 with different bit sizes.
func Integer(key string, base, size int) int64 {

	value, ok := os.LookupEnv(key)
	if ok {

		integer, err := strconv.ParseInt(value, base, size)
		if err != nil {
			return 0
		}

		return integer

	}

	return 0

}
