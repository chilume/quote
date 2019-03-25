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
	data := "Data is the new oil."
	if out := Data(); out != data {
		t.Errorf("DAta() = %q, want %q", out, data)
	}
}
