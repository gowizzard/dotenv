// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dotenv

import (
	"os"
	"strconv"
)

// Float is to get look up the environment and return is as float32 or float64.
func Float(key string, size int) float64 {

	value, ok := os.LookupEnv(key)
	if ok {

		float, err := strconv.ParseFloat(value, size)
		if err != nil {
			return 0
		}

		return float

	}

	return 0

}
