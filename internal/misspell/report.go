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

package misspell

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// CodeClimateEntry ...
type CodeClimateEntry struct {
	Description string               `json:"description"`
	Fingerprint string               `json:"fingerprint"`
	Severity    string               `json:"severity"`
	Location    CodeQualityViolation `json:"location"`
}

// CodeQualityViolation ...
type CodeQualityViolation struct {
	Path  string                       `json:"path"`
	Lines CodeQualityViolationPosition `json:"lines"`
}

// CodeQualityViolationPosition ...
type CodeQualityViolationPosition struct {
	Begin int `json:"begin"`
}

// ParseReport ...
func ParseReport(reportPath string) ([]CodeClimateEntry, error) {
	report, err := ioutil.ReadFile(reportPath)
	if err != nil {
		return []CodeClimateEntry{}, err
	}

	// Split the report by new line and process each line in turn
	lines := strings.Split(string(report), "\n")

	entries := make([]CodeClimateEntry, 0, len(lines))

	// Split on the content of the report and process each line at a time
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 4 {
			return []CodeClimateEntry{}, errors.New("unsupported misspell report line, expecting report in default format")
		}

		violationPos, _ := strconv.Atoi(parts[1])

		entries = append(entries, CodeClimateEntry{
			Description: strings.TrimSpace(parts[3]),
			Fingerprint: fmt.Sprintf("%x", md5.Sum([]byte(line))),
			Severity:    "minor",
			Location: CodeQualityViolation{
				Path: parts[0],
				Lines: CodeQualityViolationPosition{
					Begin: violationPos,
				},
			},
		})
	}

	return entries, nil
}
