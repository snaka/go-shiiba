package shiiba

import (
  "time"
)

type Activity struct {
  Date time.Time
  Count int
}

type Activities struct {
  data   []*Activity
  From  time.Time
  To    time.Time
}

// NewActivities initalize and return instance
func NewActivities() *Activities {
  now := time.Now()
  from := now.AddDate(0, 0, -364)
  a := Activities{
    data: make([]*Activity, 365),
    From: from,
    To: now,
  }
  return &a
}

//
func (p *Activities) Iterate(cb func(act *Activity) error) error {
  for i := 0; i < 10; i++ {
    err := cb(p.data[i])
    if err != nil  {
      break
    }
  }
  return nil
}
