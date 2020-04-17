package nodisk

import (
	"github.com/system-pclub/gochecker/tools/internal/lsp/foo"
)

func _() {
	foo.Foo() //@complete("F", Foo, IntFoo, StructFoo)
}
