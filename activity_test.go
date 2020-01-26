package shiiba

import (
	"reflect"
	"testing"
	"time"
)

var (
	now = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
)

func createSamples(count int) []time.Time {
	expects := make([]time.Time, count)
	for i := 0; i < count; i++ {
		offset := count - i - 1
		expects[i] = now.Add(time.Duration(24*-offset) * time.Hour)
	}
	return expects
}

func TestNewActivities(t *testing.T) {
	acts := NewActivities(now, 10)
	if acts == nil {
		t.Errorf("NewActivities is nil")
	}
	expectedTypeName := "*shiiba.Activities"
	if reflect.TypeOf(acts).String() != expectedTypeName {
		t.Errorf("fail: want %s but %T", expectedTypeName, acts)
	}
	expected, actual := time.Date(2019, 12, 22, 0, 0, 0, 0, time.UTC), acts.From
	if !expected.Equal(actual) {
		t.Errorf("fail: want %v but %v", expected, actual)
	}
	// it expects activities start from Sunday
	expectedWeekday, actualWeekday := time.Sunday, acts.First().Date.Weekday()
	if expectedWeekday != actualWeekday {
		t.Errorf("fail: want %v but %v", expectedWeekday, actualWeekday)
	}
}

func TestIterate(t *testing.T) {
	acts := NewActivities(now, 10)
	expects := createSamples(acts.Days())
	acts.Iterate(func(i int, a Activity) {
		if !expects[i].Equal(a.Date) {
			t.Errorf("#%d: date not expected: got: %v want: %v", i, a.Date, expects[i])
		}
	})
}

func TestIterateByWeekday(t *testing.T) {
	acts := NewActivities(now, 30)
	acts.IterateByWeekday(time.Sunday, func(i int, a Activity) {
		expected, actual := time.Sunday, a.Date.Weekday()
		if expected != actual {
			t.Errorf("#%d: weekday not expected. got: %v want: %v", i, actual, expected)
		}
	})
}
