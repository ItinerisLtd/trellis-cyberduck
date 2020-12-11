package lib

import (
	"errors"
	"github.com/mitchellh/cli"
	"os"
	trellisCmd "trellis-cli/cmd"
)

type Playbook struct {
	Root string
	UI   cli.Ui
}

func (p *Playbook) Run(playbookYml string, args []string) error {
	if p.Root == "" {
		return errors.New("Playbook.Root is empty; This is a flaw in the source code. Please send bug report")
	}

	if p.UI == nil {
		return errors.New("Playbook.UI is nil; This is a flaw in the source code. Please send bug report")
	}

	command := trellisCmd.CommandExecWithOutput("ansible-playbook", append([]string{playbookYml}, args...), p.UI)
	command.Dir = p.Root

	env := os.Environ()
	// To allow mockExecCommand injects its environment variables.
	if command.Env != nil {
		env = command.Env
	}
	command.Env = append(env, "ANSIBLE_RETRY_FILES_ENABLED=false")

	return command.Run()
}
