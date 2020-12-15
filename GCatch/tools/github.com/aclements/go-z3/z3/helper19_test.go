// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.9

package z3

import "testing"

func init() {
	// t.Helper was added in Go 1.9.
	tHelper = (*testing.T).Helper
}
