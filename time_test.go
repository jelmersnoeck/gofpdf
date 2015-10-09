package gofpdf_test

import (
	"testing"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func TestNow(t *testing.T) {
	clock := &gofpdf.RealClock{}

	format := func(t time.Time) string {
		return t.Format(time.RFC3339)
	}

	if format(clock.Now()) != format(time.Now()) {
		t.Errorf("Current time does not equal clock current time")
		t.FailNow()
	}

	frozenTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	clock.Freeze(frozenTime)

	if format(clock.Now()) != format(frozenTime) {
		t.Errorf("frozenTime time does not equal clock frozen time")
		t.FailNow()
	}

	clock.Unfreeze()

	if format(clock.Now()) != format(time.Now()) {
		t.Errorf("Current time does not equal clock current time")
		t.FailNow()
	}
}
