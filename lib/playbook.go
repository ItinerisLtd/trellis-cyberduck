package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Playbook struct {
	root string
	io   *Io
}

func NewPlaybook(root string, io *Io) *Playbook {
	return &Playbook{
		root: root,
		io:   io,
	}
}

func (p *Playbook) Run(name string, args []string) error {
	commandArgs := append([]string{name}, args...)

	command := exec.Command("ansible-playbook", commandArgs...)

	command.Dir = p.root

	command.Stdin = p.io.InOrStdin()
	command.Stdout = p.io.OutOrStdout()
	command.Stderr = p.io.ErrOrStderr()

	command.Env = append(os.Environ(), "ANSIBLE_RETRY_FILES_ENABLED=false")

	if _, err := fmt.Fprintf(p.io.OutOrStdout(), "Running command => ansible-playbook %s\n\n", strings.Join(commandArgs, " ")); err != nil {
		return err
	}

	return command.Run()
}
