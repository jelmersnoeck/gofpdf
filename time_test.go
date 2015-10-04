/*
 * Copyright (c) 2013-2015 Kurt Jung (Gmail: kurt.w.jung)
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

package gofpdf_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/internal/example"
	"github.com/stretchr/testify/assert"
)

func init() {
	cleanup()
}

func cleanup() {
	filepath.Walk(example.PdfDir(),
		func(path string, info os.FileInfo, err error) (reterr error) {
			if len(path) > 3 {
				if path[len(path)-4:] == ".pdf" {
					os.Remove(path)
				}
			}
			return
		})
}

func TestItem(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	tm := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	pdf.Clock().Freeze(tm)

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello World!")

	//fileStr := example.Filename("basic")
	//pdf.OutputFileAndClose(fileStr)

	buffer := &bytes.Buffer{}
	pdf.Output(buffer)
	fileBytes, _ := ioutil.ReadFile("./comparison_templates/basic.pdf")
	assert.Equal(t, buffer.Bytes(), fileBytes)
}
