// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assign_test

import (
	"testing"

	"github.com/system-pclub/GCatch/GCatch/tools/go/analysis/analysistest"
	"github.com/system-pclub/GCatch/GCatch/tools/go/analysis/passes/assign"
	"github.com/system-pclub/GCatch/GCatch/tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	tests := []string{"a"}
	if typeparams.Enabled {
		tests = append(tests, "typeparams")
	}
	analysistest.RunWithSuggestedFixes(t, testdata, assign.Analyzer, tests...)
}
