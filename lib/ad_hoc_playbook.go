package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type AdHocPlaybook struct {
	files map[string]string
	Playbook
}

func NewAdHocPlaybook(files map[string]string, io OutErrWriter) *AdHocPlaybook {
	return &AdHocPlaybook{
		files: files,
		Playbook: Playbook{
			io: io,
		},
	}
}

func (p *AdHocPlaybook) Run(path string, name string, args []string) (err error) {
	defer func() {
		if removeFilesErr := p.removeFiles(path); removeFilesErr != nil {
			err = removeFilesErr
		}
	}()

	if err := p.dumpFiles(path); err != nil {
		return err
	}

	return p.Playbook.Run(path, name, args)
}

func (p *AdHocPlaybook) dumpFiles(path string) error {
	for fileName, content := range p.files {
		destination := filepath.Join(path, fileName)
		contentByte := []byte(content)

		if err := ioutil.WriteFile(destination, contentByte, 0644); err != nil {
			return err
		}
	}

	return nil
}

func (p *AdHocPlaybook) removeFiles(path string) error {
	for fileName, _ := range p.files {
		destination := filepath.Join(path, fileName)

		if err := os.Remove(destination); err != nil {
			return err
		}
	}

	return nil
}
