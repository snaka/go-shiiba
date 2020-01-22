package shiiba

import (
  "fmt"
  "io"
  "time"
)

func ShowCalendar(out io.Writer, now time.Time, days int) error {
  acts := NewActivitiesWithFiller(now, days)
  for w := 0; w < 7; w++ {
    acts.IterateByWeekday(time.Weekday(w), func(_ int, _ Activity) error {
      fmt.Fprint(out, ".")
      return nil
    })
    fmt.Fprint(out, "\n")
  }
  return nil
}
