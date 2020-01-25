package main

import (
  "flag"
  "log"
  "os"
  "time"

  shiiba "github.com/snaka/go-shiiba"
)

var (
  showWeekday *bool = flag.Bool("w", false, "Show weekday")
  showMonth *bool = flag.Bool("m", false, "Show month")
)

func main() {
  flag.Parse()

  opts := []shiiba.Option {
    shiiba.IsShowWeekday(*showWeekday),
    shiiba.IsShowMonth(*showMonth),
  }

  err := shiiba.ShowCalendar(os.Stdout, time.Now(), 365, opts...)
  if err != nil {
    log.Fatal(err)
  }
}
