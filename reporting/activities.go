// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"sort"
	"time"
)

func ReportActivities(activities map[string]time.Duration) {
	var titles []string
	for title, _ := range activities {
		titles = append(titles, title)
	}
	sort.Strings(titles)

	for _, title := range titles {
		dur, _ := activities[title]
		fmt.Printf("%-15s %10v\n", title, dur.Truncate(time.Second))
	}
}

