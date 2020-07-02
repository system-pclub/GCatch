package shadow_test

import (
	"testing"

	"github.com/system-pclub/GCatch/tools/go/analysis/analysistest"
	"github.com/system-pclub/GCatch/tools/go/analysis/passes/shadow"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, shadow.Analyzer, "a")
}
