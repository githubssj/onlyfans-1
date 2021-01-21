package of

import (
	"fmt"
	"io"
	"os"
)

// SaveFile saves a file
func SaveFile(dir, folder, fileName string, body io.Reader) error {
	base := fmt.Sprintf("%s/%s", dir, folder)
	_, err := os.Stat(base)
	if err != nil {
		err = os.Mkdir(base, os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", base, fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, body)
	if err != nil {
		return err
	}

	return nil
}
