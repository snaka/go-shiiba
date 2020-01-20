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
func NewActivities(now time.Time) *Activities {
  from := now.AddDate(0, 0, -364)
  return &Activities{
    data: make([]*Activity, 365),
    From: from,
    To: now,
  }
}

// Iterate activities
func (p *Activities) Iterate(cb func(act *Activity) error) error {
  for i := 0; i < 10; i++ {
    err := cb(p.data[i])
    if err != nil  {
      break
    }
  }
  return nil
}
