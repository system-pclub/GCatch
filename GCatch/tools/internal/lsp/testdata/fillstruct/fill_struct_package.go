package fillstruct

import (
	h2 "net/http"

	"github.com/system-pclub/GCatch/GCatch/tools/internal/lsp/fillstruct/data"
)

func unexported() {
	a := data.B{}   //@suggestedfix("}", "refactor.rewrite")
	_ = h2.Client{} //@suggestedfix("}", "refactor.rewrite")
}
