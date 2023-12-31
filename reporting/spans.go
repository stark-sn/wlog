// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"io"
	"time"

	"s-stark.net/code/wlog/types"
)

func reportSpans(w io.Writer, day types.Day, t time.Time) {
	for _, span := range day.GetSpans(t) {
		dur := span.Duration()
		fmt.Fprintf(w, "\t%s - %s\t+ %s\n", fmtTime(span.Start), fmtTime(span.End), fmtDuration(dur))
	}
}
