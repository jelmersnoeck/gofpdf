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
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/internal/example"
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

type fontResourceType struct {
}

func (f fontResourceType) Open(name string) (rdr io.Reader, err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(example.FontFile(name))
	if err == nil {
		rdr = bytes.NewReader(buf)
		fmt.Printf("Generalized font loader reading %s\n", name)
	}
	return
}

// Convert 'ABCDEFG' to, for example, 'A,BCD,EFG'
func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}

func lorem() string {
	return "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod " +
		"tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis " +
		"nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis " +
		"aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat " +
		"nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui " +
		"officia deserunt mollit anim id est laborum."
}

// This example demonstrates the generation of a simple PDF document. Note that
// since only core fonts are used (in this case Arial, a synonym for
// Helvetica), an empty string can be specified for the font directory in the
// call to New(). Note also that the example.Filename() and example.Summary()
// functions belong to a separate, internal package and are not part of the
// gofpdf library. If an error occurs at some point during the construction of
// the document, subsequent method calls exit immediately and the error is
// finally retreived with the output call where it can be handled by the
// application.
func Example() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.SetMargins(0, 0, 0)
	//CellFormat(w, h float64, txtStr string, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string)
	//pdf.CellFormat(50, 50, "This is a very long string that should be cut across multiple lines", "1", 1, "CT", false, 0, "")
	pdf.TextBox(50, 50, "This is a very long string that should be cut across multiple lines don't you think? What is your opinion on this whole text box thing?", "1", "CM")
	fileStr := example.Filename("basic")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/basic.pdf
}
