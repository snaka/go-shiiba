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

// NewActivities initalize and return instance
func NewActivities(now time.Time, days int) *Activities {
  from := now.Add(time.Duration(-(days - 1) * 24) * time.Hour)
  a := &Activities{
    data: make([]*Activity, days),
    From: from,
    To: now,
  }
  a.init()
  return a
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
