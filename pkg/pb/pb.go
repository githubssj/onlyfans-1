package pb

import (
	pbar "github.com/schollz/progressbar/v3"
)

func NewProgressBar(size int64) *pbar.ProgressBar {
	return pbar.DefaultBytes(
		size,
		"downloading onlyfans media...",
	)
}
