// Package workmate: JSON configuration loading component to manage Go workspace
// Provides simple configuration file parsing of workspace and project paths
// Works great with go-work package workspace structures
//
// workmate: Go 工作区管理的 JSON 配置加载组件
// 提供工作区和项目路径的简单配置文件解析
// 与 go-work 包的工作区结构良好配合
package workmate

import (
	"encoding/json"
	"os"

	"github.com/go-mate/go-work/workspace"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

// Workspace represents a JSON configuration structure to manage workspace
// Contains workspace root DIR and list of project paths supporting multi-project development
// Designed to support simple JSON serialization and deserialization
//
// Workspace 代表工作区管理的 JSON 配置结构
// 包含工作区根目录和用于多项目开发的项目路径列表
// 设计支持便捷的 JSON 序列化和反序列化
type Workspace struct {
	WorkRoot string   `json:"workRoot"` // Workspace root DIR // 工作区根目录
	Projects []string `json:"projects"` // Project paths in this workspace // 此工作区中的项目路径
}

// NewWorkspace loads workspace configuration from JSON file at specified path
// Opens and parses JSON configuration file into Workspace struct
// Returns configured Workspace prepared to use with go-work operations
//
// NewWorkspace 从指定路径的 JSON 文件加载工作区配置
// 打开并解析 JSON 配置文件到 Workspace 结构
// 返回配置好的工作区，可用在 go-work 操作中
func NewWorkspace(path string) *Workspace {
	file := rese.P1(os.Open(path))
	defer rese.F0(file.Close)
	res := &Workspace{}
	must.Done(json.NewDecoder(file).Decode(res))
	return res
}

// GetWorkspace converts workmate.Workspace to go-work workspace.Workspace format
// Creates compatible workspace structure to use with go-work package operations
// Enables seamless integration between configuration loading and workspace execution
//
// GetWorkspace 将 workmate.Workspace 转换为 go-work workspace.Workspace 格式
// 创建兼容的工作区结构，用在 go-work 包操作中
// 实现配置加载和工作区执行之间的无缝集成
func (wc *Workspace) GetWorkspace() *workspace.Workspace {
	return &workspace.Workspace{
		WorkRoot: wc.WorkRoot,
		Projects: wc.Projects,
	}
}
