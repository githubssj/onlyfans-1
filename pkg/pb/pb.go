package pb

import (
	pbar "github.com/schollz/progressbar/v3"
)

func NewProgressBar() *pbar.ProgressBar {
	return pbar.DefaultBytes(
		-1,
		"downloading onlyfans media...",
	)
}
