package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type AdHocPlaybook struct {
	files map[string]string
	root  string
	ui    *Ui
}

func NewAdHocPlaybook(files map[string]string, root string, ui *Ui) *AdHocPlaybook {
	return &AdHocPlaybook{
		files: files,
		root:  root,
		ui:    ui,
	}
}

func (p *AdHocPlaybook) Run(name string, args []string) (err error) {
	defer func() {
		if removeFilesErr := p.removeFiles(); removeFilesErr != nil {
			err = removeFilesErr
		}
	}()

	if err := p.dumpFiles(); err != nil {
		return err
	}

	playbook := NewPlaybook(p.root, p.ui)
	return playbook.Run(name, args)
}

func (p *AdHocPlaybook) dumpFiles() error {
	for fileName, content := range p.files {
		destination := filepath.Join(p.root, fileName)
		contentByte := []byte(content)

		if err := ioutil.WriteFile(destination, contentByte, 0644); err != nil {
			return err
		}
	}

	return nil
}

func (p *AdHocPlaybook) removeFiles() error {
	for fileName, _ := range p.files {
		destination := filepath.Join(p.root, fileName)

		if err := os.Remove(destination); err != nil {
			return err
		}
	}

	return nil
}
