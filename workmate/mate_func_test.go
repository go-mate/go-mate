package workmate_test

import (
	"testing"

	"github.com/go-mate/go-mate/workmate"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

func TestNewMateFunc(t *testing.T) {
	mateFunc := workmate.NewMateFunc(func(path string) {
		res := string(rese.V1(osexec.ExecInPath(path, "ls", "-a")))
		t.Log(res)
	})

	path := runpath.PARENT.Path()

	must.Done(mateFunc(path))
}
