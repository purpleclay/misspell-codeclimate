/*
Copyright (c) 2022 Purple Clay

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package misspell_test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/purpleclay/misspell-codeclimate/internal/misspell"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:misspell
func TestParseReport(t *testing.T) {
	line1 := `docs/README.md:1:14: "incorectly" is a misspelling of "incorrectly"`
	line2 := `docs/test.txt:1:14: "incorectly" is a misspelling of "incorrectly"`
	path := writeFile(t, fmt.Sprintf("%s\n%s", line1, line2))

	report, err := misspell.ParseReport(path)
	require.NoError(t, err)

	require.Len(t, report, 2)
	assert.Equal(t, `"incorectly" is a misspelling of "incorrectly"`, report[0].Description)
	assert.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(line1))), report[0].Fingerprint)
	assert.Equal(t, "minor", report[0].Severity)
	assert.Equal(t, "docs/README.md", report[0].Location.Path)
	assert.Equal(t, 1, report[0].Location.Lines.Begin)

	assert.Equal(t, `"incorectly" is a misspelling of "incorrectly"`, report[1].Description)
	assert.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(line2))), report[1].Fingerprint)
	assert.Equal(t, "minor", report[1].Severity)
	assert.Equal(t, "docs/test.txt", report[1].Location.Path)
	assert.Equal(t, 1, report[1].Location.Lines.Begin)
}

func writeFile(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	f := filepath.Join(dir, "misspell-report.txt")

	err := ioutil.WriteFile(f, []byte(content), 0o644)
	require.NoError(t, err)

	return f
}

func TestParseReport_ReportNotFound(t *testing.T) {
	_, err := misspell.ParseReport("unknown-report.txt")

	assert.EqualError(t, err, "open unknown-report.txt: no such file or directory")
}

func TestParseReport_UnsupportedReportFormat(t *testing.T) {
	path := writeFile(t, "unexpected format")

	_, err := misspell.ParseReport(path)
	assert.EqualError(t, err, "unsupported misspell report line, expecting report in default format. Received: unexpected format")
}
