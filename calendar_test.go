package shiiba

import (
  "testing"
  "time"
  "reflect"
)

func TestNewActivities(t *testing.T) {
  now, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z00:00")
  a := NewActivities(now)
  if a == nil {
    t.Errorf("NewActivities is nil")
  }
  expectedTypeName := "*shiiba.Activities"
  if reflect.TypeOf(a).String() != expectedTypeName {
     t.Errorf("fail: want %s but %T", expectedTypeName, a)
  }
  expectedFromDate, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z00:00")
  if !expectedFromDate.Equal(a.From) {
    t.Errorf("fail: want %v but %v", expectedFromDate, a.From)
  }
}
