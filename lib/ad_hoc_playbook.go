package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type AdHocPlaybook struct {
	Playbook
	Files map[string]string
}

func (p *AdHocPlaybook) Run(playbookYml string, args []string) (err error) {
	if len(p.Files) == 0 {
		return errors.New("AdHocPlaybook.Files is empty; This is a flaw in the source code. Please send bug report")
	}

	defer func() {
		if removeFilesErr := p.removeFiles(); removeFilesErr != nil {
			err = removeFilesErr
		}
	}()

	if err := p.dumpFiles(); err != nil {
		return err
	}

	return p.Playbook.Run(playbookYml, args)
}

func (p *AdHocPlaybook) dumpFiles() error {
	for fileName, content := range p.Files {
		destination := filepath.Join(p.Root, fileName)
		contentByte := []byte(content)

		if err := ioutil.WriteFile(destination, contentByte, 0644); err != nil {
			return err
		}
	}

	return nil
}

func (p *AdHocPlaybook) removeFiles() error {
	for fileName, _ := range p.Files {
		destination := filepath.Join(p.Root, fileName)

		if err := os.Remove(destination); err != nil {
			return err
		}
	}

	return nil
}
