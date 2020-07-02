package badstmt

import (
	"github.com/system-pclub/GCatch/tools/internal/lsp/foo"
)

func _() {
	defer foo.F //@complete("F", Foo, IntFoo, StructFoo),diag(" //", "LSP", "function must be invoked in defer statement")
	go foo.F //@complete("F", Foo, IntFoo, StructFoo)
}
