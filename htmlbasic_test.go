package gofpdf_test

import (
	"strconv"
	"testing"

	"github.com/jung-kurt/gofpdf"
)

func TestGetStringWidth(t *testing.T) {
	var width float64

	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.SetFont("Helvetica", "", 20)

	html := pdf.HTMLBasicNew()
	width = html.GetStringWidth("This is a string")

	if width != 128.92 {
		t.Errorf("Width should be `128.92`, is `" + floatToStr(width) + "`")
	}
}

func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}
