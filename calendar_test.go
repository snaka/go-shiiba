package shiiba

import (
  "bytes"
  "testing"
  "time"
)

func TestShowCalendar(t *testing.T) {
  buf := &bytes.Buffer{}
  now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
  ShowCalendar(buf, now, 30)
  output := buf.String()

  expected := `.....
.....
.....
.....
....
....
....
`
  if expected != output {
    t.Errorf("fail: calendar output invalid want %v but %v", expected, output)
  }
}
