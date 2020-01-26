package shiiba

import (
	"time"
)

// Activity stores activity count at date
type Activity struct {
	Date  time.Time
	Count int
}

// Activities stores collection of activities during a specified period
type Activities struct {
	data []*Activity
	From time.Time
	To   time.Time
}

func dateBefore(base time.Time, before int) time.Time {
	return base.Add(time.Duration(-(before)*24) * time.Hour)
}

// NewActivitiesWithFiller initialize with filler that adjust data starts from Sunday
func NewActivities(now time.Time, days int) *Activities {
  offset := int(dateBefore(now, days-1).Weekday())
	adjustedDays := days + offset
	from := dateBefore(now, adjustedDays-1)
	a := &Activities{
		data: make([]*Activity, adjustedDays),
		From: from,
		To:   now,
	}
	a.init()
	return a
}

// init activities
func (p *Activities) init() {
	length := len(p.data)
	for i := 0; i < length; i++ {
		p.data[i] = &Activity{Date: p.From.Add(time.Duration(i*24) * time.Hour)}
	}
}

// Days returns data count
func (p *Activities) Days() int {
  return len(p.data)
}


// Iterate activities
func (p *Activities) Iterate(cb func(int, Activity)) {
	for i, v := range p.data {
		cb(i, *v)
	}
}

// IterateByWeekday iterates activities by each weekday
func (p *Activities) IterateByWeekday(weekday time.Weekday, cb func(int, Activity)) {
	col := 0
	p.Iterate(func(i int, a Activity) {
		if a.Date.Weekday() == weekday {
			cb(col, a)
			col++
		}
	})
}

// First returns first activity of data
func (p *Activities) First() Activity {
	return *p.data[0]
}
