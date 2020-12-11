package lib

import (
	"errors"
	"os"
	"os/exec"
)

type Playbook struct {
	Root string
}

func (p *Playbook) Run(playbookYml string, args []string) error {
	if p.Root == "" {
		return errors.New("Playbook.Root is empty; This is a flaw in the source code. Please send bug report")
	}

	commandArgs := append([]string{playbookYml}, args...)

	command := exec.Command("ansible-playbook", commandArgs...)
	command.Dir = p.Root
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	env := os.Environ()
	// To allow mockExecCommand injects its environment variables.
	if command.Env != nil {
		env = command.Env
	}
	command.Env = append(env, "ANSIBLE_RETRY_FILES_ENABLED=false")

	return command.Run()
}
