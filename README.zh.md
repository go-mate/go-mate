[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-mate/go-mate/release.yml?branch=main&label=BUILD)](https://github.com/go-mate/go-mate/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-mate/go-mate)](https://pkg.go.dev/github.com/go-mate/go-mate)
[![Coverage Status](https://img.shields.io/coveralls/github/go-mate/go-mate/main.svg)](https://coveralls.io/github/go-mate/go-mate?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://github.com/go-mate/go-mate)
[![GitHub Release](https://img.shields.io/github/release/go-mate/go-mate.svg)](https://github.com/go-mate/go-mate/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mate/go-mate)](https://goreportcard.com/report/github.com/go-mate/go-mate)

# go-mate

Go 项目开发的配置管理助手组件。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 安装

```bash
go get github.com/go-mate/go-mate/workmate
```

## 使用方法

### 基础用法

```go
package main

import (
    "github.com/go-mate/go-mate/workmate"
)

func main() {
    // 加载配置
    config := workmate.NewWorkspace("mate.json")

    // 转换为 go-work 格式
    workspace := config.GetWorkspace()

    // 现在可以用 workspace 配合 go-work 命令使用
}
```


### 配置文件

创建配置 `mate.json`：

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

## mate 生态系统

```bash
# 根据需要获取工具
go install github.com/go-mate/depbump/cmd/depbump@latest   # 依赖管理
go install github.com/go-mate/go-commit/cmd/go-commit@latest # 带 Go 格式化的 Git 提交
go install github.com/go-mate/go-lint/cmd/go-lint@latest   # 代码检查
go install github.com/go-mate/go-work/cmd/go-work@latest   # 工作区管理
go install github.com/go-mate/tago/cmd/tago@latest         # Git 标签管理
```

go-mate 在 mate 生态系统中扮演**配置管理组件**的角色：

```
配置管理
    ↓
go-mate (JSON 配置加载)
    ↓
go-work (工作区执行)
    ↓
开发工具 (depbump, go-commit, go-lint, tago)
```

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/go-mate/go-mate.svg?variant=adaptive)](https://starchart.cc/go-mate/go-mate)