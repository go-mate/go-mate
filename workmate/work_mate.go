package workmate

import (
	"path/filepath"

	"github.com/yyle88/erero"
	"github.com/yyle88/osexistpath/osmustexist"
)

type WorksMates []*WorkMate

func (mates WorksMates) Execute() error {
	for _, mate := range mates {
		if err := mate.Execute(); err != nil {
			return erero.Wro(err)
		}
	}
	return nil
}

type WorkMate struct {
	WorkRoot string // Absolute path of the directory containing go.work

	BeforeWorks []MateFunc

	Projects []*ProjectMate // List of subprojects in the workspace

	FinishWorks []MateFunc
}

func (w *WorkMate) WithMoreBeforeWork(mateFunc MateFunc) *WorkMate {
	w.BeforeWorks = append(w.BeforeWorks, mateFunc)
	return w
}

func (w *WorkMate) WithMoreFinishWork(mateFunc MateFunc) *WorkMate {
	w.FinishWorks = append(w.FinishWorks, mateFunc)
	return w
}

func (w *WorkMate) Execute() error {
	if w.WorkRoot != "" {
		osmustexist.MustRoot(w.WorkRoot)
		osmustexist.MustFile(filepath.Join(w.WorkRoot, "go.work"))

		for _, run := range w.BeforeWorks {
			if err := run(w.WorkRoot); err != nil {
				return erero.Wro(err)
			}
		}
	}
	for _, project := range w.Projects {
		if err := project.Execute(); err != nil {
			return erero.Wro(err)
		}
	}
	if w.WorkRoot != "" {
		for _, run := range w.FinishWorks {
			if err := run(w.WorkRoot); err != nil {
				return erero.Wro(err)
			}
		}
	}
	return nil
}

func (w *WorkMate) GetProjectPaths() []string {
	paths := make([]string, 0, len(w.Projects))
	for _, project := range w.Projects {
		paths = append(paths, project.GetPath())
	}
	return paths
}

type ProjectMate struct {
	path string
	runs []MateFunc
}

func NewProjectMate(path string, runs []MateFunc) *ProjectMate {
	return &ProjectMate{
		path: path,
		runs: runs,
	}
}

func (c *ProjectMate) WithMoreFunc(mateFunc MateFunc) *ProjectMate {
	c.runs = append(c.runs, mateFunc)
	return c
}

func (c *ProjectMate) Execute() error {
	path := c.GetPath()
	for _, run := range c.runs {
		if err := run(path); err != nil {
			return erero.Wro(err)
		}
	}
	return nil
}

func (c *ProjectMate) GetPath() string {
	osmustexist.MustRoot(c.path)
	osmustexist.MustFile(filepath.Join(c.path, "go.mod"))
	return c.path
}
