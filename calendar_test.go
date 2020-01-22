package shiiba

import (
  "testing"
  "time"
  "reflect"
)

var (
  now = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
)

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
  a := NewActivities(now, 3)
  expectFirst := now
  expects := []time.Time {
    expectFirst.Add(time.Duration(24 * -2) * time.Hour),
    expectFirst.Add(time.Duration(24 * -1) * time.Hour),
    expectFirst,
  }
  a.Iterate(func(i int, act Activity) error {
    expected := expects[i]
    actual := act.Date
    if expected.Equal(actual) == false {
      t.Errorf("#%d: date not expected: got: %v want: %v", i, act.Date, expects[i])
    }
    return nil
  })
}
