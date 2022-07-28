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
	"encoding/json"
	"fmt"
	"io"

	"github.com/purpleclay/misspell-codeclimate/internal/misspell"
	"github.com/spf13/cobra"
)

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
			report, err := misspell.ParseReport(opt.ReportFile)
			if err != nil {
				return err
			}

			data, err := json.Marshal(report)
			if err != nil {
				return err
			}

			fmt.Fprintln(out, string(data))
			return nil
		},
	}

	rootCmd.Flags().StringVar(&opt.ReportFile, "file", "", "Path to the misspell report to parse")
	rootCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(newVersionCmd(out))
	rootCmd.AddCommand(newManPagesCmd(out))
	rootCmd.AddCommand(newCompletionCmd(out))

	return rootCmd.ExecuteContext(ctx.Background())
}
