package mateconfig

import (
	"encoding/json"
	"os"

	"github.com/go-mate/go-work/workconfig"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

type WorkspaceConfig struct {
	WorkRoot string   `json:"workRoot"` // 工作区根目录
	Projects []string `json:"projects"` // 该 Workspace 中的项目路径
}

func NewWorkspaceConfig(path string) *WorkspaceConfig {
	file := rese.P1(os.Open(path))
	defer func() {
		must.Done(file.Close())
	}()
	res := &WorkspaceConfig{}
	must.Done(json.NewDecoder(file).Decode(res))
	return res
}

func (wc *WorkspaceConfig) GetWorkspace() *workconfig.Workspace {
	return &workconfig.Workspace{
		WorkRoot: wc.WorkRoot,
		Projects: wc.Projects,
	}
}
