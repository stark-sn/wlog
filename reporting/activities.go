// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"io"
	"sort"
	"time"
)

func reportActivities(w io.Writer, activities map[string]time.Duration) {
	var titles []string
	for title, _ := range activities {
		titles = append(titles, title)
	}
	sort.Strings(titles)

	for _, title := range titles {
		dur, _ := activities[title]
		fmt.Fprintf(w, "\t%s\t+ %v\n", title, fmtDuration(dur))
	}
}
