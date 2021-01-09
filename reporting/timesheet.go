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
	var actWeek time.Duration
	acts := make(map[string]act)

	w := tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)

	for _, date := range dates {
		day, _ := week.Days[date]
		spanTime := sumWorkingTimeDay(day, t)
		breakTime := sumBreakTime(day, t)

		dayTime := spanTime - breakTime
		var actDay time.Duration

		dayActs := make(map[string]act)
		for _, activity := range day.GetActivities(t) {
			dur := activity.End.Sub(activity.Start)
			actDay += dur
			sumActs(acts, activity.Title, dur)
			sumActs(dayActs, activity.Title, dur)
		}

		actWeek += actDay
		dayActs["Untracked"] = act{dur: dayTime - actDay}

		fmt.Fprintf(w, "%s\t%s\n", date, fmtDuration(dayTime))
		printActs(w, dayActs, nil, "")

		weekTime += dayTime
		fmt.Fprintln(w, "\t")
	}

	acts["Untracked"] = act{dur: weekTime - actWeek}

	fmt.Fprintf(w, "Week\t%s\n", fmtDuration(weekTime))
	printActs(w, acts, nil, "")
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
	} else {
		a.taskDur += dur
		a.isTask = true
	}

	acts[splits[0]] = a
}

func printActs(w io.Writer, acts map[string]act, parent *act, padding string) {

	var titles []string
	for title, _ := range acts {
		titles = append(titles, title)
	}
	sort.Strings(titles)

	for i, title := range titles {
		act := acts[title]

		elementPadding, childPadding := getPadding(padding, i == len(titles)-1, act.isTask, parent != nil && parent.isTask)

		fmt.Fprintf(w, "%s %s\t%s\n", elementPadding, title, fmtDuration(act.dur))

		if act.isTask && len(act.sub) > 0 {
			fmt.Fprintf(w, "%s╟─ ...\t%s\n", childPadding, fmtDuration(act.taskDur))
		}

		printActs(w, act.sub, &act, childPadding)
	}
}

func getPadding(padding string, isLast bool, isTask bool, isSubTask bool) (string, string) {

	elementPadding := padding
	childPadding := padding

	if isLast {
		childPadding += "   "

		if isSubTask {
			if isTask {
				elementPadding += "╚"
			} else {
				elementPadding += "╙"
			}
		} else {
			if isTask {
				elementPadding += "╘"
			} else {
				elementPadding += "└"
			}
		}
	} else {
		if isSubTask {
			childPadding += "║  "

			if isTask {
				elementPadding += "╠"
			} else {
				elementPadding += "╟"
			}
		} else {
			childPadding += "│  "

			if isTask {
				elementPadding += "╞"
			} else {
				elementPadding += "├"
			}
		}
	}

	if isTask {
		elementPadding += "═"
	} else {
		elementPadding += "─"
	}

	return elementPadding, childPadding
}

type act struct {
	dur     time.Duration
	isTask  bool
	taskDur time.Duration
	sub     map[string]act
}
