package lib

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAdHocPlaybook_Run(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	playbook := NewPlaybook(bufferIo)
	err := playbook.Run(".", "ad_hoc_playbook.yml", []string{"foo", "bar"})

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	want := "Received command: ansible-playbook ad_hoc_playbook.yml foo bar"
	if got := bufferIo.OutString(); !strings.Contains(got, want) {
		t.Errorf("want %s, got %s", want, got)
	}
}
