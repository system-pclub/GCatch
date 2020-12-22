// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"fmt"
	"regexp"
	"testing"
)

var tHelper func(t *testing.T)

func wantPanic(t *testing.T, match string, cb func()) {
	if tHelper != nil {
		tHelper(t)
	}

	re, err := regexp.Compile(match)
	if err != nil {
		t.Fatal("error compiling regexp: ", err)
	}
	defer func() {
		if tHelper != nil {
			tHelper(t)
		}
		err := recover()
		if err == nil {
			t.Fatalf("want panic matching %q, got success", match)
		}
		s := fmt.Sprint(err)
		if !re.MatchString(s) {
			t.Fatalf("want panic matching %q, got %s", match, s)
		}
	}()
	cb()
}

func simplifyBool(t *testing.T, ctx *Context, x Bool) bool {
	if tHelper != nil {
		tHelper(t)
	}

	y := ctx.Simplify(x, nil).(Bool)
	eq, ok := y.AsBool()
	if !ok {
		fmt.Println("Simplify(A) = B, want bool literal. A:", x, " B:", y)
		t.Fatal()
	}
	return eq
}
