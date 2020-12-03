// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"io"
	"s-stark.net/code/wlog/types"
	"time"
)

func reportSpans(w io.Writer, day types.Day, t time.Time) {
	for _, span := range day.GetSpans(t) {
		dur := span.End.Sub(span.Start).Round(time.Second)
		fmt.Fprintf(w, "\t%s - %s\t+ %v\n", fmtTime(span.Start), fmtTime(span.End), fmtDuration(dur))
	}
}
