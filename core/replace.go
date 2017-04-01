package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Replacer struct {
	Name        string
	Expression  string
	Replacement string
	Nop         bool
}

func (o Replacer) Replace(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match(o.Name, fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		old := string(read)
		new := strings.Replace(old, o.Expression, o.Replacement, -1)
		if old != new {
			if !o.Nop {
				fmt.Println("Replace in: " + path)
				err = ioutil.WriteFile(path, []byte(new), 0)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println("(NOP) Replace in: " + path)
			}
		} else {
			fmt.Println("No match for replace expression: " + path)
		}
	} else {
		fmt.Println("No match for file name: " + path)
	}

	return nil
}
