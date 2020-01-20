package shiiba

import (
  "testing"
  "reflect"
)

func TestNewActivities(t *testing.T) {
  a := NewActivities()
  if a == nil {
    t.Errorf("NewActivities is nil")
  }
  expectedTypeName := "*shiiba.Activities"
  if reflect.TypeOf(a).String() != expectedTypeName {
     t.Errorf("fail: want %s but %T", expectedTypeName, a)
  }
}
