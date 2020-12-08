// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"io"
	"os"
	"s-stark.net/code/wlog/types"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

// Display a timesheet report
func Timesheet(week types.Week, t time.Time) error {
	if week.Days == nil {
		fmt.Println("You're were not in this week.")
		return nil
	}

	var dates []string
	for date, _ := range week.Days {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	var weekTime time.Duration
	acts := make(map[string]act)

	w := tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)

	for _, date := range dates {
		day, _ := week.Days[date]
		spanTime := sumWorkingTimeDay(day, t)
		breakTime := sumBreakTime(day, t)

		dayTime := spanTime - breakTime

		dayActs := make(map[string]act)
		for _, activity := range day.GetActivities(t) {
			dur := activity.End.Sub(activity.Start)
			sumActs(acts, activity.Title, dur)
			sumActs(dayActs, activity.Title, dur)
		}

		fmt.Fprintf(w, "%s\t%s\n", date, fmtDuration(dayTime))
		printActs(w, dayActs, "")

		weekTime += dayTime
		fmt.Fprintln(w, "\t")
	}

	fmt.Fprintf(w, "Week\t%s\n", fmtDuration(weekTime))
	printActs(w, acts, "")
	w.Flush()

	return nil
}

func sumActs(acts map[string]act, title string, dur time.Duration) {
	splits := strings.SplitN(title, ":", 2)

	a, _ := acts[splits[0]]
	a.dur += dur

	if len(splits) > 1 {
		if acts[splits[0]].sub == nil {
			a.sub = make(map[string]act)
		}
		sumActs(a.sub, splits[1], dur)
	}

	acts[splits[0]] = a
}

func printActs(w io.Writer, acts map[string]act, padding string) {

	var titles []string
	for title, _ := range acts {
		titles = append(titles, title)
	}
	sort.Strings(titles)

	for i, title := range titles {
		act := acts[title]

		isLast := i == len(titles)-1

		myPadding := padding
		childPadding := padding

		if isLast {
			myPadding += "└─"
			childPadding += "   "
		} else {
			myPadding += "├─"
			childPadding += "│  "
		}

		fmt.Fprintf(w, "%s %s\t%s\n", myPadding, title, fmtDuration(act.dur))

		printActs(w, act.sub, childPadding)
	}
}

type act struct {
	title string
	dur   time.Duration
	sub   map[string]act
}
