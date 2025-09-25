[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-mate/go-mate/release.yml?branch=main&label=BUILD)](https://github.com/go-mate/go-mate/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-mate/go-mate)](https://pkg.go.dev/github.com/go-mate/go-mate)
[![Coverage Status](https://img.shields.io/coveralls/github/go-mate/go-mate/main.svg)](https://coveralls.io/github/go-mate/go-mate?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-mate/go-mate.svg)](https://github.com/go-mate/go-mate/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mate/go-mate)](https://goreportcard.com/report/github.com/go-mate/go-mate)

# go-mate

Configuration management assistant component that helps develop golang projects.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

📋 **JSON Configuration Loading**: Simple and efficient JSON configuration file parsing
🔄 **Workspace Integration**: Perfect integration with go-work package workspace management
⚡ **Lightweight Design**: Focused on core configuration management with minimum dependencies
🌍 **Structure Conversion**: Convenient conversion from configuration to workspace structures
📝 **Standard Format**: Support of standard JSON configuration format serialization and deserialization

## Installation

```bash
go get github.com/go-mate/go-mate/workmate
```

## Usage

### JSON Configuration File Format

Create configuration file `mate.json`:

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

### Basic Usage

```go
package main

import (
    "github.com/go-mate/go-mate/workmate"
    "github.com/go-mate/go-work/worksexec"
    "github.com/yyle88/osexec"
)

func main() {
    // Load configuration file
    workspace := workmate.NewWorkspace("mate.json")

    // Convert to go-work format
    goWorkspace := workspace.GetWorkspace()

    // Create execution configuration
    execConfig := osexec.NewCommandConfig()
    workspaces := []*workspace.Workspace{goWorkspace}
    config := worksexec.NewWorksExec(execConfig, workspaces)

    // Use configuration to execute commands
    // config.ForeachWorkRun(...)
}
```

### Configuration Structure

- **workRoot**: Workspace root DIR path
- **projects**: List of project paths included in this workspace

### Integration with go-work

The main purpose of go-mate is to provide configuration file support when using the go-work package:

```go
// Load configuration from JSON
workspace := workmate.NewWorkspace("config.json")

// Convert to go-work compatible format
goWorkspace := workspace.GetWorkspace()

// Use with go-work operations
config := worksexec.NewWorksExec(execConfig, []*workspace.Workspace{goWorkspace})
```

## API Documentation

### Workspace Struct

```go
type Workspace struct {
    WorkRoot string   `json:"workRoot"` // Workspace root DIR
    Projects []string `json:"projects"` // Project paths in this workspace
}
```

### Main Methods

- `NewWorkspace(path string) *Workspace`: Load workspace configuration from JSON file
- `GetWorkspace() *workspace.Workspace`: Convert to go-work package compatible format

## Use Cases

### Multi-project Development Environment
- Unified management of multiple Go project configurations
- Batch execution of cross-project commands
- Workspace-stage package management

### Automation Scripts
- Configuration-driven project management
- Support of CI/CD pipeline configurations
- Simplified operations with complex project structures

### Development Solutions Integration
- Collaboration with the mate ecosystem components
- Project configuration of IDEs and editors
- Support of rapid development environment switching

## Design Approach

### Simple Design
- Focused API design
- Focus on core configuration management functions
- Simple-to-understand and use interfaces

### Integration
- Perfect integration with go-work package
- Support of mate ecosystem collaboration
- Standardized configuration format

### Scalable Design
- Flexible JSON configuration structure
- Support of future feature extensions
- Compatible-first design

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
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
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/go-mate/go-mate.svg?variant=adaptive)](https://starchart.cc/go-mate/go-mate)