package gofpdf_test

import (
	"testing"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func TestNow(t *testing.T) {
	clock := &gofpdf.RealClock{}

	if clock.Now().Format(time.RFC3339) != time.Now().Format(time.RFC3339) {
		t.Errorf("Current time does not equal clock current time")
		t.FailNow()
	}

	frozenTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	clock.Freeze(frozenTime)

	if clock.Now().Format(time.RFC3339) != frozenTime.Format(time.RFC3339) {
		t.Errorf("frozenTime time does not equal clock frozen time")
		t.FailNow()
	}

	clock.Unfreeze()

	if clock.Now().Format(time.RFC3339) != time.Now().Format(time.RFC3339) {
		t.Errorf("Current time does not equal clock current time")
		t.FailNow()
	}
}
