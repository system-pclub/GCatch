// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgfact_test

import (
	"testing"

	"github.com/system-pclub/gochecker/tools/go/analysis/analysistest"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/pkgfact"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, pkgfact.Analyzer, "c")
}
