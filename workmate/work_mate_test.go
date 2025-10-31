package workmate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

// TestNewWorkspace verifies workspace configuration loading from JSON
// Creates temp DIR, writes test config, loads workspace, and validates content
// Environment setup uses rese/must while test verification uses require
//
// TestNewWorkspace 验证从 JSON 加载工作区配置
// 创建临时 DIR，写入测试配置，加载工作区并验证内容
// 环境设置使用 rese/must，测试验证使用 require
func TestNewWorkspace(t *testing.T) {
	// Get temp DIR during test - auto-creates and auto-cleans
	// 获取测试中的临时 DIR - 自动创建和清理
	tempDIR := t.TempDir()

	// Create test configuration
	// 创建测试配置
	configPath := filepath.Join(tempDIR, "test-config.json")
	config := &Workspace{
		WorkRoot: "/Users/test/projects",
		Projects: []string{
			"/Users/test/projects/project1",
			"/Users/test/projects/project2",
		},
	}

	// Write configuration to JSON - must succeed during test setup
	// 将配置写入 JSON - 测试设置必须成功
	fp := rese.P1(os.Create(configPath))
	must.Done(json.NewEncoder(fp).Encode(config))
	rese.F0(fp.Close)

	// Load workspace from configuration
	// 从配置加载工作区
	workspace := NewWorkspace(configPath)

	// Check loaded workspace - test verification uses require
	// 验证加载的工作区 - 测试验证使用 require
	require.NotNil(t, workspace)
	require.Equal(t, "/Users/test/projects", workspace.WorkRoot)
	require.Len(t, workspace.Projects, 2)
	require.Equal(t, "/Users/test/projects/project1", workspace.Projects[0])
	require.Equal(t, "/Users/test/projects/project2", workspace.Projects[1])
}
