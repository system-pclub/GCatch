package util

import (
	"fmt"
	"runtime"
)

// Debugfln prints a debug information with a new line. It attaches a prefix `[DEBUG]` to the format string
// It returns the number of bytes written and any write error encountered.
func Debugfln(format string, a ...interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	progLoc := ""
	if ok {
		progLoc = runtime.FuncForPC(pc).Name()
	} else {
		progLoc = "(unknown)"
	}
	fmt.Printf("[DEBUG] "+progLoc+": "+format+"\n", a...)
}
