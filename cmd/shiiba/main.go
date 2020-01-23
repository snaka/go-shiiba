package main

import (
  "log"
  "os"
  "time"

  shiiba "github.com/snaka/go-shiiba"
)

func main() {
  err := shiiba.ShowCalendar(os.Stdout, time.Now(), 365, shiiba.IsShowWeekday(true))
  if err != nil {
    log.Fatal(err)
  }
}
