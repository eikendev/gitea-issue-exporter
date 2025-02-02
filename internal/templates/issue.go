package templates

const IssueTmpl = `# {{.Issue.Title}}

**Created by {{.Issue.Poster.UserName}} on {{.Issue.Created.Format "2006-01-02 15:04:05"}}**

Labels: {{range .Issue.Labels}}{{.Name}} {{end}}

{{.Issue.Body}}

{{ if .CommentList }}
## Comments ({{len .CommentList}})

{{ range .CommentList }}
**{{.Poster.UserName}} commented on {{.Created.Format "2006-01-02 15:04:05"}}:**

{{.Body}}

---

{{ end }}
{{- end }}
`
