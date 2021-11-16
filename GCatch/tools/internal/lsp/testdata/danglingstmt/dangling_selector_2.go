package danglingstmt

import "github.com/system-pclub/GCatch/GCatch/tools/internal/lsp/foo"

func _() {
	foo. //@rank(" //", Foo)
	var _ = []string{foo.} //@rank("}", Foo)
}
