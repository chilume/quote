// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quote collects sayings about Data.
package quote

import "math/rand"

var sayings = []string{
	"Data is the new oil.",
	"In God we trust, all others bring data.",
	"If we have data, let's look at data. If all we have are opinions, let's go with mine",
}

var size int

func init() {
	size = len(sayings)
}

// Data returns a random quote on data.
func Data() string {
	return sayings[rand.Intn(size)]
}
