// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"io"
	"sort"
	"time"
)

func reportNonWorkingTime(w io.Writer, nonWorkingTime map[string]time.Duration) {
	var titles []string
	for title, _ := range nonWorkingTime {
		titles = append(titles, title)
	}
	sort.Strings(titles)

	for _, title := range titles {
		dur, _ := nonWorkingTime[title]
		fmt.Fprintf(w, "\t%s\t+ %s\n", title, fmtDuration(dur))
	}
}
