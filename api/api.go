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
package api

import (
	"context"
	"fmt"
	"github.com/google/go-github/v59/github"
	"github.com/tnagatomi/gh-mrlabel/option"
	"strings"
)

func CreateLabel(client *github.Client, label option.Label, repo option.Repo) error {
	githubLabel := &github.Label{
		Name:        github.String(label.Name),
		Description: github.String(label.Description),
		Color:       github.String(label.Color),
	}
	_, _, err := client.Issues.CreateLabel(context.Background(), repo.Owner, repo.Repo, githubLabel)

	if err != nil {
		if strings.Contains(err.Error(), "already_exists") {
			return fmt.Errorf("label %q already exists for repository %q", label, repo)
		}
		return err
	}

	return nil
}
