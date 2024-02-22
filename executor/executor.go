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
package executor

import (
	"fmt"
	"github.com/google/go-github/v59/github"
	"github.com/tnagatomi/gh-mrlabel/api"
	"github.com/tnagatomi/gh-mrlabel/parser"
	"io"
	"net/http"
)

// Executor composites github.Client and has dry-run option
type Executor struct {
	client *github.Client
	dryRun bool
}

// NewExecutor returns new Executor
func NewExecutor(client *http.Client, dryrun bool) (*Executor, error) {
	return &Executor{
		client: github.NewClient(client),
		dryRun: dryrun,
	}, nil
}

// Create creates labels across multiple repositories
func (e *Executor) Create(out io.Writer, repoOption string, labelOption string) error {
	labels, err := parser.Label(labelOption)
	if err != nil {
		return fmt.Errorf("failed to parse label option: %v", err)
	}
	repos, err := parser.Repo(repoOption)
	if err != nil {
		return fmt.Errorf("failed to parse repo option: %v", err)
	}

	for _, repo := range repos {
		for _, label := range labels {
			if e.dryRun {
				fmt.Fprintf(out, "Would create label %q for repository %q\n", label, repo)
				continue
			}

			err = api.CreateLabel(e.client, label, repo)
			if err != nil {
				fmt.Fprintf(out, "Failed to create label %q for repository %q: %v\n", label, repo, err)
				continue
			}
			fmt.Fprintf(out, "Created label %q for repository %q\n", label, repo)
		}
	}

	return nil
}
