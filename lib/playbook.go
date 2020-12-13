package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var execCommand = exec.Command

type Playbook struct {
	io OutErrWriter
}

func NewPlaybook(io OutErrWriter) *Playbook {
	return &Playbook{
		io: io,
	}
}

func (p *Playbook) Run(path string, name string, args []string) error {
	commandArgs := append([]string{name}, args...)

	command := execCommand("ansible-playbook", commandArgs...)

	command.Dir = path

	command.Stdout = p.io.OutOrStdout()
	command.Stderr = p.io.ErrOrStderr()

	// To allow execCommand injects its environment variables.
	env := os.Environ()
	if command.Env != nil {
		env = command.Env
	}
	command.Env = append(env, "ANSIBLE_RETRY_FILES_ENABLED=false")

	if _, err := fmt.Fprintf(p.io.OutOrStdout(), "Running command => ansible-playbook %s\n\n", strings.Join(commandArgs, " ")); err != nil {
		return err
	}

	return command.Run()
}
