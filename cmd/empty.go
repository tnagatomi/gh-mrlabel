/*
Copyright © 2024 Takayuki Nagatomi <tommyt6073@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/tnagatomi/gh-mrlabel/executor"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// NewEmptyCmd represents the empty command
func NewEmptyCmd(in io.Reader, out io.Writer) *cobra.Command {
	var emptyCmd = &cobra.Command{
		Use:   "empty",
		Short: "Empties labels across multiple repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := api.NewHTTPClient(api.ClientOptions{})
			if err != nil {
				return fmt.Errorf("failed to create gh http client: %v", err)
			}

			e, err := executor.NewExecutor(client, dryRun)
			if err != nil {
				return fmt.Errorf("failed to create exector: %v", err)
			}

			if !dryRun {
				confirmed, err := confirm(in, out)
				if err != nil {
					return fmt.Errorf("failed to confirm execution: %v", err)
				}
				if !confirmed {
					fmt.Fprintf(out, "Canceled execution\n")
					return nil
				}
			}

			err = e.Empty(out, repos)
			if err != nil {
				return fmt.Errorf("failed to empty labels: %v", err)
			}

			return nil
		},
	}

	return emptyCmd
}

func init() {
	emptyCmd := NewEmptyCmd(os.Stdin, os.Stdout)
	rootCmd.AddCommand(emptyCmd)
}
