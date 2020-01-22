package shiiba

import (
  "testing"
  "time"
  "reflect"
)

var (
  now = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
)

func createSamples(count int) []time.Time {
  expects := make([]time.Time, count)
  for i := 0; i < count; i++ {
    offset := count - i - 1
    expects[i] = now.Add(time.Duration(24 * -offset) * time.Hour)
  }
  return expects
}

func TestNewActivities(t *testing.T) {
  a := NewActivities(now, 10)
  if a == nil {
    t.Errorf("NewActivities is nil")
  }
  expectedTypeName := "*shiiba.Activities"
  if reflect.TypeOf(a).String() != expectedTypeName {
     t.Errorf("fail: want %s but %T", expectedTypeName, a)
  }
  expectedFromDate := time.Date(2019, 12, 23, 0, 0, 0, 0, time.UTC)
  if !expectedFromDate.Equal(a.From) {
    t.Errorf("fail: want %v but %v", expectedFromDate, a.From)
  }
}

func TestIterate(t *testing.T) {
  expects := createSamples(10)
  a := NewActivities(now, len(expects))
  a.Iterate(func(i int, act Activity) error {
    if !expects[i].Equal(act.Date) {
      t.Errorf("#%d: date not expected: got: %v want: %v", i, act.Date, expects[i])
    }
    return nil
  })
}
