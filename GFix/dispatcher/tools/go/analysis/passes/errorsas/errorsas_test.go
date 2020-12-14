// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.13

package errorsas_test

import (
	"testing"

	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/analysistest"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/passes/errorsas"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, errorsas.Analyzer, "a")
}
