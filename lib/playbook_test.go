package lib

import (
	"os/exec"
	"strings"
	"testing"
)

var bufferIo = NewBufferIo()

func TestPlaybook_Run(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	playbook := NewPlaybook(bufferIo)
	err := playbook.Run(".", "playbook.yml", []string{"foo", "bar"})

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	want := "Received command: ansible-playbook playbook.yml foo bar"
	if got := bufferIo.OutString(); !strings.Contains(got, want) {
		t.Errorf("want %s, got %s", want, got)
	}
}
