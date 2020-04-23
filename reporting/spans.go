// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"io"
	"time"
	"s-stark.net/code/wlog/types"
)

func reportSpans(w io.Writer, day types.Day) {
	for _, span := range day.Spans {
		dur := span.End.Sub(span.Start).Round(time.Second)
		fmt.Fprintf(w, "\t%s - %s\t+ %v\n", fmtTime(span.Start), fmtTime(span.End), fmtDuration(dur))
	}
}

