// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"fmt"
	"regexp"
	"testing"
)

func expectPanic(t *testing.T, pattern string, f func()) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatal("bad regexp: ", err)
	}
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("want panic matching %q, got none", pattern)
			} else if s := fmt.Sprint(err); !re.MatchString(s) {
				t.Fatalf("want panic matching %q, got %s", pattern, err)
			}
	}()
	f()
}

func TestErrorHandling(t *testing.T) {
	ctx := NewContext(nil)
	x := ctx.BVConst("x", 1)
	y := ctx.BVConst("y", 2)
	expectPanic(t, "are incompatible", func() { x.Eq(y) })
}
