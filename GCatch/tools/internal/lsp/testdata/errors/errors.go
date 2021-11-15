package errors

import (
	"github.com/system-pclub/GCatch/GCatch/tools/internal/lsp/types"
)

func _() {
	bob.Bob() //@complete(".")
	types.b //@complete(" //", Bob_interface)
}
