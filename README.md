[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-mate/go-mate/release.yml?branch=main&label=BUILD)](https://github.com/go-mate/go-mate/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-mate/go-mate)](https://pkg.go.dev/github.com/go-mate/go-mate)
[![Coverage Status](https://img.shields.io/coveralls/github/go-mate/go-mate/main.svg)](https://coveralls.io/github/go-mate/go-mate?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://github.com/go-mate/go-mate)
[![GitHub Release](https://img.shields.io/github/release/go-mate/go-mate.svg)](https://github.com/go-mate/go-mate/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mate/go-mate)](https://goreportcard.com/report/github.com/go-mate/go-mate)

# go-mate

Configuration management assistant component that helps develop golang projects.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Installation

```bash
go get github.com/go-mate/go-mate/workmate
```

## Usage

### Basic Usage

```go
package main

import (
    "github.com/go-mate/go-mate/workmate"
)

func main() {
    // Load configuration
    config := workmate.NewWorkspace("mate.json")

    // Convert to go-work format
    workspace := config.GetWorkspace()

    // Now use workspace with go-work commands
}
```


### Configuration

Create configuration `mate.json`:

```json
{
  "workRoot": "/Users/admin/go-projects/mate",
  "projects": [
    "/Users/admin/go-projects/mate/go-work",
    "/Users/admin/go-projects/mate/depbump",
    "/Users/admin/go-projects/mate/tago",
    "/Users/admin/go-projects/mate/go-lint",
    "/Users/admin/go-projects/mate/go-mate"
  ]
}
```

## mate Ecosystem

```bash
# Get tools as needed
go install github.com/go-mate/depbump/cmd/depbump@latest   # Deps management
go install github.com/go-mate/go-commit/cmd/go-commit@latest # Git commits with Go formatting
go install github.com/go-mate/go-lint/cmd/go-lint@latest   # Code linting
go install github.com/go-mate/go-work/cmd/go-work@latest   # Workspace management
go install github.com/go-mate/tago/cmd/tago@latest         # Git tag management
```

go-mate serves as the **configuration management component** in the mate ecosystem:

```
Configuration
       ↓
   go-mate (JSON config loading)
       ↓
   go-work (workspace execution)
       ↓
Development Tools (depbump, go-commit, go-lint, tago)
```

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/go-mate/go-mate.svg?variant=adaptive)](https://starchart.cc/go-mate/go-mate)