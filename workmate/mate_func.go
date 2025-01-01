package workmate

type MateFunc = func(path string) error

func NewMateFunc(run func(path string)) MateFunc {
	return func(path string) error {
		run(path)
		return nil
	}
}
