// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loopclosure_test

import (
	"github.com/system-pclub/GCatch/GCatch/tools/internal/typeparams"
	"testing"

	"github.com/system-pclub/GCatch/GCatch/tools/go/analysis/analysistest"
	"github.com/system-pclub/GCatch/GCatch/tools/go/analysis/passes/loopclosure"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	tests := []string{"a", "golang.org/..."}
	if typeparams.Enabled {
		tests = append(tests, "typeparams")
	}
	analysistest.Run(t, testdata, loopclosure.Analyzer, tests...)
}
