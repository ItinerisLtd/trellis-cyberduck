package lib

import (
	"os"
	"os/exec"
)

type Playbook struct {
	root string
	ui   *Ui
}

func NewPlaybook(root string, ui *Ui) *Playbook {
	return &Playbook{
		root: root,
		ui:   ui,
	}
}

func (p *Playbook) Run(name string, args []string) error {
	commandArgs := append([]string{name}, args...)

	command := exec.Command("ansible-playbook", commandArgs...)

	command.Dir = p.root

	command.Stdin = p.ui.InOrStdin()
	command.Stdout = p.ui.OutOrStdout()
	command.Stderr = p.ui.ErrOrStderr()

	command.Env = append(os.Environ(), "ANSIBLE_RETRY_FILES_ENABLED=false")

	return command.Run()
}
