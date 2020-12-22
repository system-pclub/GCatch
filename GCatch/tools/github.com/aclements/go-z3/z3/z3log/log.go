// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package z3log exposes Z3's interaction log.
//
// The interaction log is a low-level trace of all Z3 API calls.
package z3log

import "unsafe"

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

// Open creates a Z3 interaction log in a file called filename.
//
// It returns false if it fails to open the log.
func Open(filename string) bool {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	//return C.Z3_open_log(cfilename) != 0
	return C.Z3_open_log(cfilename) == true
}

// Append emits text to the Z3 interaction log.
func Append(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.Z3_append_log(ctext)
}

// Close closes the Z3 interaction log file.
func Close() {
	C.Z3_close_log()
}
