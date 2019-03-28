// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quote

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("LC_ALL", "en")
}

func TestData(t *testing.T) {
	out := Data()
	for _, s := range sayings {
		if s == out {
			return
		}
	}
	t.Errorf("Data() = %v, want %v", out, sayings)
}
