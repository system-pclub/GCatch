package nodisk

import (
	"github.com/system-pclub/GCatch/tools/internal/lsp/foo"
)

func _() {
	foo.Foo() //@complete("F", Foo, IntFoo, StructFoo)
}
