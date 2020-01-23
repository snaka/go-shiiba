package shiiba

import (
  "fmt"
  "io"
  "time"
)

// Configs represents calendar behavior
type Configs struct {
  IsShowWeekday bool
  IsShowMonth bool
}

// Option sets configure
type Option func(*Configs)

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
  args := &Configs {}
  for _, option := range options {
    option(args)
  }

  acts := NewActivitiesWithFiller(now, days)
  for w := 0; w < 7; w++ {
    if args.IsShowWeekday {
      var s string
      if w % 2 == 0 {
        s = "   "
      } else {
        s = shortNameOfWeekday(w)
      }
      fmt.Fprintf(out, "%s ", s)
    }
    acts.IterateByWeekday(time.Weekday(w), func(_ int, _ Activity) error {
      fmt.Fprint(out, ".")
      return nil
    })
    fmt.Fprint(out, "\n")
  }
  return nil
}

func shortNameOfWeekday(w int) string {
  return time.Weekday(w).String()[:3]
}
