package cyberduck

import (
	"fmt"
	"strings"
	"time"

	"github.com/itinerisltd/trellis-cyberduck/lib"
)

const cyberduckOpenYml = `
---
- name: 'Trellis Cyberduck: Open SFTP connetions via Cyberduck'
  hosts: web:&{{ env }}
  gather_facts: false
  connection: local
  tasks:
    - name: Print debug message
      debug:
        msg: "Generating bookmark file at {{ dest }}"
    - name: Include deploy role variables
      include_vars: 
        dir: "{{ playbook_dir }}/roles/deploy/defaults"
    - name: Include trellis-cyberduck variables
      include_vars:
        file: "{{ playbook_dir }}/cyberduck_defaults.yml"
    - name: Generate bookmark file
      template:
        src: cyberduck_bookmark.j2
        dest: "{{ dest }}"
        mode: '0600'
      with_dict: "{{ wordpress_sites }}"
      when: item.key == site
    - name: Open SFTP connection
      command: "open {{ dest }}" 
    - name: Cleanup bookmark file
      file:
        path: "{{ dest }}"
        state: absent
`

const cyberduckDefaultsYml = `
---
project_uploads_path: "{{ project_root }}/shared/uploads"
project_current_symlink_path: "{{ project_root + '/' + project_current_path }}"
`

const cyberduckBookmarkJ2 = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Protocol</key>
	<string>sftp</string>
	<key>Hostname</key>
	<string>{{ ansible_host }}</string>
	<key>Port</key>
	<string>{{ ansible_port | default('22') }}</string>
	<key>Username</key>
	<string>{{ lookup('vars', user + '_user') }}</string>
	<key>Path</key>
	<string>{{ lookup('vars', directory) }}</string>
</dict>
</plist>
`

type Opener struct {
	io lib.OutErrWriter
}

func NewOpener() *Opener {
	return &Opener{
		io: lib.NewIo(),
	}
}

func (o *Opener) SetIo(io lib.OutErrWriter) {
	o.io = io
}

func (o *Opener) Open(path string, environment string, siteName string, user string, directory string) error {
	playbook := lib.NewAdHocPlaybook(map[string]string{
		"cyberduck_open.yml":    strings.TrimSpace(cyberduckOpenYml) + "\n",
		"cyberduck_defaults.yml": strings.TrimSpace(cyberduckDefaultsYml) + "\n",
		"cyberduck_bookmark.j2": strings.TrimSpace(cyberduckBookmarkJ2) + "\n",
	}, o.io)

	playbookArgs := []string{
		"-e", "dest=" + fmt.Sprintf("%s/cyberduck-%d.duck", path, time.Now().UnixNano()),
		"-e", "env=" + environment,
		"-e", "site=" + siteName,
		"-e", "user=" + user,
		"-e", "directory=" + directory,
	}

	return playbook.Run(path, "cyberduck_open.yml", playbookArgs)
}
