package links

import (
	"fmt" //@link(re`".*"`,"https://godoc.org/fmt")

	"github.com/system-pclub/GCatch/GCatch/tools/internal/lsp/foo" //@link(re`".*"`,"https://godoc.org/github.com/system-pclub/GCatch/GCatch/tools/internal/lsp/foo")
)

var (
	_ fmt.Formatter
	_ foo.StructFoo
)
