package shiiba

import (
  "time"
)

// Activity stores activity count at date
type Activity struct {
  Date time.Time
  Count int
}

// Activities stores collection of activities during a specified period
type Activities struct {
  data   []*Activity
  From  time.Time
  To    time.Time
}

func dateBefore(base time.Time, before int) time.Time {
  return base.Add(time.Duration(-(before) * 24) * time.Hour)
}

// NewActivities initalize and return instance
func NewActivities(now time.Time, days int) *Activities {
  from := dateBefore(now, days - 1)
  a := &Activities{
    data: make([]*Activity, days),
    From: from,
    To: now,
  }
  a.init()
  return a
}

// NewActivitiesWithFiller initialize with filler that adjust data starts from Sunday
func NewActivitiesWithFiller(now time.Time, days int) *Activities {
  adjustDays := int(dateBefore(now, days - 1).Weekday())
  adjustedDays := days + adjustDays
  return NewActivities(now, adjustedDays)
}

// init activities
func (p *Activities) init() {
  length := len(p.data)
  for i := 0; i < length; i++ {
    p.data[i] = &Activity { Date: p.From.Add(time.Duration(i * 24) * time.Hour) }
  }
}

// Iterate activities
func (p *Activities) Iterate(cb func(int, Activity) error) error {
  for i, v := range p.data {
    err := cb(i, *v)
    if err != nil  {
      break
    }
  }
  return nil
}

// IterateByWeekday iterates activities by each weekday
func (p *Activities) IterateByWeekday(weekday time.Weekday, cb func(int, Activity) error) error {
  p.Iterate(func(i int, a Activity) error {
    if a.Date.Weekday() == weekday {
      cb(i, a)
    }
    return nil
  })
  return nil
}

// First returns first activity of data
func (p *Activities) First() Activity {
  return *p.data[0]
}
