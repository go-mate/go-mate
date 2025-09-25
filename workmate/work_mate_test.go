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

func TestNewWorkspace(t *testing.T) {
	// Get temporary DIR for test
	// 获取测试用的临时 DIR
	tempDIR := t.TempDir()

	// Create test configuration file
	// 创建测试配置文件
	configPath := filepath.Join(tempDIR, "test-config.json")
	config := &Workspace{
		WorkRoot: "/Users/test/projects",
		Projects: []string{
			"/Users/test/projects/project1",
			"/Users/test/projects/project2",
		},
	}

	// Write configuration to file
	// 将配置写入文件
	file := rese.P1(os.Create(configPath))
	must.Done(json.NewEncoder(file).Encode(config))
	rese.F0(file.Close)

	// Load workspace from configuration file
	// 从配置文件加载工作区
	workspace := NewWorkspace(configPath)

	// Verify loaded workspace
	// 验证加载的工作区
	require.NotNil(t, workspace)
	require.Equal(t, "/Users/test/projects", workspace.WorkRoot)
	require.Len(t, workspace.Projects, 2)
	require.Equal(t, "/Users/test/projects/project1", workspace.Projects[0])
	require.Equal(t, "/Users/test/projects/project2", workspace.Projects[1])
}
