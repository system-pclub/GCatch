package c

import "github.com/system-pclub/GCatch/GCatch/tools/internal/lsp/rename/b"

func _() {
	b.Hello() //@rename("Hello", "Goodbye")
}
