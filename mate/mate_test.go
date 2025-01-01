package mate

import (
	"testing"

	"github.com/go-mate/go-mate/workmate"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

func TestMate_NewWorkMate(t *testing.T) {
	path := runpath.PARENT.Up(1)

	mate := NewMate("", []string{path})

	workMate := mate.NewWorkMate(workmate.NewMateFunc(func(path string) {
		t.Log(string(rese.V1(osexec.ExecInPath(path, "ls", "-a"))))
	}))

	require.NoError(t, workMate.Execute())
}
