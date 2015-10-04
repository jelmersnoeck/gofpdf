package gofpdf

import "time"

type Clock interface {
	Now() time.Time
	Freeze(time.Time)
}

type RealClock struct {
	frozen     bool
	frozenTime time.Time
}

func (c *RealClock) Now() time.Time {
	if c.frozen {
		return c.frozenTime
	} else {
		return time.Now()
	}
}

func (c *RealClock) Freeze(frozenTime time.Time) {
	c.frozen = true
	c.frozenTime = frozenTime
}
