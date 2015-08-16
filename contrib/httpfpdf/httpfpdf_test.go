package httpfpdf_test

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpfpdf"
	"path/filepath"
)

const (
	cnGofpdfDir  = "./"
	cnExampleDir = cnGofpdfDir + "/pdf"
)

func exampleFilename(baseStr string) string {
	return filepath.Join(cnExampleDir, baseStr+".pdf")
}

func summary(err error, fileStr string) {
	if err == nil {
		fileStr = filepath.ToSlash(fileStr)
		fmt.Printf("Successfully generated %s\n", fileStr)
	} else {
		fmt.Println(err)
	}
}

func ExampleFpdf_AddRemoteImage() {
	pdf := gofpdf.New("", "", "", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()

	pdf.Text(20, 20, "This is text")

	url := "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/UPC-A-036000291452.png/220px-UPC-A-036000291452.png"
	httpfpdf.RegisterRemoteImage(pdf, url, "")
	pdf.Image(url, 100, 100, 20, 20, false, "", 0, "")

	fileStr := exampleFilename("Fpdf_AddRemoteImage")
	err := pdf.OutputFileAndClose(fileStr)
	summary(err, fileStr)
	// Output:
	// Successfully generated pdf/Fpdf_AddRemoteImage.pdf
}
