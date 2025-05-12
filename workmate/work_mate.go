package workmate

import (
	"encoding/json"
	"os"

	"github.com/go-mate/go-work/workspace"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

type Workspace struct {
	WorkRoot string   `json:"workRoot"` // 工作区根目录
	Projects []string `json:"projects"` // 该 Workspace 中的项目路径
}

func NewWorkspace(path string) *Workspace {
	file := rese.P1(os.Open(path))
	defer rese.F0(file.Close)
	res := &Workspace{}
	must.Done(json.NewDecoder(file).Decode(res))
	return res
}

func (wc *Workspace) GetWorkspace() *workspace.Workspace {
	return &workspace.Workspace{
		WorkRoot: wc.WorkRoot,
		Projects: wc.Projects,
	}
}
