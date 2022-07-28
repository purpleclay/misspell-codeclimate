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

package cmd

import (
	ctx "context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type codeClimateEntry struct {
	Description string               `json:"description"`
	Fingerprint string               `json:"fingerprint"`
	Severity    string               `json:"severity"`
	Location    codeQualityViolation `json:"location"`
}

type codeQualityViolation struct {
	Path  string                       `json:"path"`
	Lines codeQualityViolationPosition `json:"lines"`
}

type codeQualityViolationPosition struct {
	Begin int `json:"begin"`
}

type options struct {
	ReportFile string
}

func Execute(out io.Writer) error {
	opt := options{}

	rootCmd := &cobra.Command{
		Use:          "misspell-codeclimate",
		Short:        "Turn that misspell report into a GitLab compatible codeclimate report",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return transformReport(opt, out)
		},
	}

	rootCmd.Flags().StringVar(&opt.ReportFile, "f", "", "Path to the misspell report to parse")

	rootCmd.AddCommand(newVersionCmd(out))
	rootCmd.AddCommand(newManPagesCmd(out))
	rootCmd.AddCommand(newCompletionCmd(out))

	return rootCmd.ExecuteContext(ctx.Background())
}

func transformReport(opt options, out io.Writer) error {
	report, err := ioutil.ReadFile(opt.ReportFile)
	if err != nil {
		return err
	}

	// Split the report by new line and process each line in turn
	lines := strings.Split(string(report), "\n")

	entries := make([]codeClimateEntry, 0, len(lines))

	// Split on the content of the report and process each line at a time
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 4 {
			return errors.New("unsupported misspell report line, expecting report in default format")
		}

		violationPos, _ := strconv.Atoi(parts[1])

		entries = append(entries, codeClimateEntry{
			Description: parts[3],
			Fingerprint: fmt.Sprintf("%x", md5.Sum([]byte(line))),
			Severity:    "minor",
			Location: codeQualityViolation{
				Path: parts[0],
				Lines: codeQualityViolationPosition{
					Begin: violationPos,
				},
			},
		})
	}

	data, err := json.Marshal(entries)
	if err != nil {
		return err
	}

	fmt.Fprintln(out, string(data))
	return nil
}
