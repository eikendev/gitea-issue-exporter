# 📦 Gitea Issue Exporter

A command-line tool to export issues from a Gitea repository to markdown files.

## ✨ Features

- 📝 Exports issues to individual markdown files
- 💬 Includes issue comments
- 🏷️ Preserves issue metadata (labels, author, timestamps)
- 🔍 Filters by issue state (open, closed, or all)
- ⚡ Excludes pull requests
- 📚 Supports pagination for large repositories

## 🚀 Installation

```bash
go install github.com/eikendev/gitea-issue-exporter@latest
```

## 🎮 Usage

```bash
gitea-issue-exporter [flags] owner/repo
```

### 🎯 Flags

- `-url string`: Gitea instance URL (default "https://gitea.com")
- `-token string`: Gitea access token (required)
- `-output string`: Output directory for markdown files (default "issues")
- `-state string`: Issue state to export (all, open, closed) (default "all")

### 💡 Example

```bash
export GITEA_TOKEN="your-token-here"
gitea-issue-exporter -token $GITEA_TOKEN -output ./issues octocat/hello-world
```

### 📄 Output Format

Each issue is exported to a separate markdown file named `issue-{number}.md` containing:

- 📌 Issue title
- 👤 Author and creation timestamp
- 🏷️ Labels
- 📝 Issue body
- 💭 Comments with author and timestamp information

## 🔑 Authentication

You need a Gitea access token to use this tool. Generate one from your Gitea instance:

1. ⚙️ Go to Settings -> Applications
2. 🔧 Under "Manage Access Tokens", create a token
3. 🔒 Save the token securely
