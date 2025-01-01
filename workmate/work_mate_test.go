package workmate_test

import (
	"testing"

	"github.com/go-mate/go-mate/workmate"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

func TestNewProjectConfig(t *testing.T) {
	path := runpath.PARENT.Up(1)
	t.Log(path)

	config := workmate.NewProjectMate(path, []workmate.MateFunc{
		workmate.NewMateFunc(func(path string) {
			t.Log(string(rese.V1(osexec.ExecInPath(path, "ls", "-a"))))
		}),
		workmate.NewMateFunc(func(path string) {
			t.Log(string(rese.V1(osexec.ExecInPath(path, "go", "mod", "tidy", "-e"))))
		}),
	})

	require.NoError(t, config.Execute())
}
