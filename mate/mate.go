package mate

import "github.com/go-mate/go-mate/workmate"

type Mate struct {
	WorkRoot string
	Projects []string
}

func NewMate(workRoot string, projects []string) *Mate {
	return &Mate{
		WorkRoot: workRoot,
		Projects: projects,
	}
}

func (mate *Mate) NewWorkMate(mateFunc workmate.MateFunc) *workmate.WorkMate {
	projects := make([]*workmate.ProjectMate, 0, len(mate.Projects))

	for _, path := range mate.Projects {
		projects = append(projects, workmate.NewProjectMate(path, []workmate.MateFunc{mateFunc}))
	}

	return &workmate.WorkMate{
		WorkRoot:    mate.WorkRoot,
		BeforeWorks: nil,
		Projects:    projects,
		FinishWorks: nil,
	}
}
