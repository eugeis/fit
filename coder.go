package fit

import (
	"io"
	"os"
	"golang.org/x/text/encoding/charmap"
	"github.com/tarent/logrus"
)

func AnsiToUtf8(filePath string, targetPath string, nop bool) (err error) {
	var file, target *os.File
	if file, err = os.Open(filePath); err == nil {
		if target, err = os.Create(targetPath); err == nil {
			if !nop {
				logrus.Infof("Encode the '%v' to the '%v'", file.Name(), target.Name())
				r := charmap.ISO8859_1.NewDecoder().Reader(file)
				io.Copy(target, r)
			} else {
				logrus.Infof("Skip the '%v' to the '%v'", file.Name(), target.Name())
			}
			target.Close()
		}
		file.Close()
	}
	return
}
