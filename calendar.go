package shiiba

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

// Configs represents calendar behavior
type Configs struct {
	ActivityProvider Provider
	IsShowWeekday    bool
	IsShowMonth      bool
}

// Option sets configure
type Option func(*Configs)

// ActivityProvider provides activity data like Qiita, Github, etc.
func ActivityProvider(service string) Option {
	return func(args *Configs) {
		args.ActivityProvider = getProvider(service)
	}
}

// IsShowWeekday indicates calendar shows weekday or not
func IsShowWeekday(on bool) Option {
	return func(args *Configs) {
		args.IsShowWeekday = on
	}
}

// IsShowMonth indicates calendar shows month or not
func IsShowMonth(on bool) Option {
	return func(args *Configs) {
		args.IsShowMonth = on
	}
}

// ShowCalendar puts calendar to 'out' buffer
func ShowCalendar(out io.Writer, now time.Time, days int, options ...Option) error {
	args := &Configs{
		ActivityProvider: NullProvider,
		IsShowWeekday:    false,
		IsShowMonth:      false,
	}
	for _, option := range options {
		option(args)
	}

	acts := NewActivitiesWithFiller(now, days)
	args.ActivityProvider(acts)

	if args.IsShowMonth {
		if args.IsShowWeekday {
			fmt.Fprint(out, "    ")
		}
		printMonths(out, acts)
	}
	for w := 0; w < 7; w++ {
		if args.IsShowWeekday {
			var s string
			if w%2 == 0 {
				s = "   "
			} else {
				s = shortNameOfWeekday(w)
			}
			fmt.Fprintf(out, "%s ", s)
		}
		acts.IterateByWeekday(time.Weekday(w), func(_ int, a Activity) {
			if a.Count > 0 {
				fmt.Fprint(out, "o")
			} else {
				fmt.Fprint(out, ".")
			}
		})
		fmt.Fprint(out, "\n")
	}
	return nil
}

func printMonths(out io.Writer, acts *Activities) {
	var buf bytes.Buffer
	iterateMonths(acts, func(col, month int) {
		if buf.Len() < col {
			buf.WriteString(strings.Repeat(" ", col-buf.Len()))
		}
		buf.WriteString(shortMonthName(month))
	})
	buf.WriteString("\n")
	buf.WriteTo(out)
}

func iterateMonths(acts *Activities, cb func(col, month int)) {
	var prev time.Month
	acts.IterateByWeekday(time.Sunday, func(col int, a Activity) {
		if remainDays := daysToNextMonth(a.Date); remainDays < 21 {
			return
		}
		if current := a.Date.Month(); current != prev {
			cb(col, int(current))
			prev = current
		}
	})
}

func daysToNextMonth(base time.Time) int {
	y, m, d := base.Date()
	currentMonth := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	nextMonth := time.Date(y, time.Month(m+1), 1, 0, 0, 0, 0, time.UTC)
	remain := nextMonth.Sub(currentMonth)
	return int(remain.Hours() / 24)
}

func shortMonthName(m int) string {
	return time.Month(m).String()[:3]
}

func shortNameOfWeekday(w int) string {
	return time.Weekday(w).String()[:3]
}
