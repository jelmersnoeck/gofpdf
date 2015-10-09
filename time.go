/*
 * Copyright (c) 2013-2014 Kurt Jung (Gmail: kurt.w.jung)
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gofpdf

import "time"

// A Clock represents a moment in time. This moment in time is represented in
// nanoseconds (just as time.Time) but can also be frozen to a specific moment
// in time.
type Clock interface {
	Now() time.Time
	Freeze(time.Time)
	Unfreeze()
}

// A RealClock is an implementation of Clock which deals with remembering the
// potentially frozen time.
type RealClock struct {
	frozen     bool
	frozenTime time.Time
}

// Now gives back the time of the Clock. If the Clock has previously been
// frozen, it will return the frozen time. If not, it will use the current
// available time.
func (c *RealClock) Now() time.Time {
	if c.frozen {
		return c.frozenTime
	} else {
		return time.Now()
	}
}

// Freeze will freeze the clock at a specific given time.
func (c *RealClock) Freeze(frozenTime time.Time) {
	c.frozen = true
	c.frozenTime = frozenTime
}

// Unfreeze will unfreeze the clock, meaning that all future calls will use the
// available current time.
func (c *RealClock) Unfreeze() {
	c.frozen = false
}
