package export

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/eikendev/gitea-issue-exporter/internal/templates"

	"code.gitea.io/sdk/gitea"
)

type Exporter struct {
	client *gitea.Client
	tmpl   *template.Template
	outDir string
}

func New(client *gitea.Client, outDir string) (*Exporter, error) {
	tmpl, err := template.New("issue").Parse(templates.IssueTmpl)
	if err != nil {
		return nil, err
	}

	return &Exporter{
		client: client,
		tmpl:   tmpl,
		outDir: outDir,
	}, nil
}

func (e *Exporter) getAllComments(owner, repo string, index int64) ([]*gitea.Comment, error) {
	page := 1
	pageSize := 30
	var allComments []*gitea.Comment

	for {
		comments, _, err := e.client.ListIssueComments(owner, repo, index, gitea.ListIssueCommentOptions{
			ListOptions: gitea.ListOptions{
				Page:     page,
				PageSize: pageSize,
			},
		})
		if err != nil {
			return nil, err
		}

		allComments = append(allComments, comments...)

		if len(comments) < pageSize {
			break
		}
		page++
	}

	return allComments, nil
}

func (e *Exporter) Export(owner, repo string, state gitea.StateType) error {
	if err := os.MkdirAll(e.outDir, 0o750); err != nil {
		return err
	}

	issues, _, err := e.client.ListRepoIssues(owner, repo, gitea.ListIssueOption{
		State: state,
		Type:  gitea.IssueTypeIssue, // Only real issues, no PRs
	})
	if err != nil {
		return err
	}

	for _, issue := range issues {
		comments, err := e.getAllComments(owner, repo, issue.Index)
		if err != nil {
			return err
		}

		data := struct {
			Issue       *gitea.Issue
			CommentList []*gitea.Comment
		}{
			Issue:       issue,
			CommentList: comments,
		}

		filename := filepath.Join(e.outDir, fmt.Sprintf("issue-%d.md", issue.Index))
		f, err := os.Create(filename) // #nosec G304
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filename, err)
		}
		defer func() {
			if cerr := f.Close(); cerr != nil {
				err = fmt.Errorf("failed to close file %s: %w", filename, cerr)
			}
		}()

		if err := e.tmpl.Execute(f, data); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	}

	return nil
}
