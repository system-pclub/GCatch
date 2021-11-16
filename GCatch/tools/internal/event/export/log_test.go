// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package export_test

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/system-pclub/GCatch/GCatch/tools/internal/event"
	"github.com/system-pclub/GCatch/GCatch/tools/internal/event/core"
	"github.com/system-pclub/GCatch/GCatch/tools/internal/event/export"
	"github.com/system-pclub/GCatch/GCatch/tools/internal/event/keys"
	"github.com/system-pclub/GCatch/GCatch/tools/internal/event/label"
)

func ExampleLog() {
	ctx := context.Background()
	event.SetExporter(timeFixer(export.LogWriter(os.Stdout, false)))
	anInt := keys.NewInt("myInt", "an integer")
	aString := keys.NewString("myString", "a string")
	event.Log(ctx, "my event", anInt.Of(6))
	event.Error(ctx, "error event", errors.New("an error"), aString.Of("some string value"))
	// Output:
	// 2020/03/05 14:27:48 my event
	// 	myInt=6
	// 2020/03/05 14:27:48 error event: an error
	// 	myString="some string value"
}

func timeFixer(output event.Exporter) event.Exporter {
	at, _ := time.Parse(time.RFC3339Nano, "2020-03-05T14:27:48Z")
	return func(ctx context.Context, ev core.Event, lm label.Map) context.Context {
		copy := core.CloneEvent(ev, at)
		return output(ctx, copy, lm)
	}
}
