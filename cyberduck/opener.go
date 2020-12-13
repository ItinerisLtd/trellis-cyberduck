package cyberduck

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/itinerisltd/trellis-cyberduck/lib"
)

const cyberduckOpenYml = `
---
- name: 'Trellis CLI: Dump database credentials'
  hosts: web:&{{ env }}
  gather_facts: false
  connection: local
  tasks:
    - name: Print debug message
      debug:
        msg: "Generating bookmark file at {{ dest }}"
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
	<string>{{ is_admin | bool | ternary(admin_user, web_user) }}</string>
	<key>Path</key>
	<string>{{ project_root | default(www_root + '/' + item.key) | regex_replace('^~\/','') }}/{{ item.current_path | default('current') }}</string>
</dict>
</plist>
`

type Opener struct {
	io *lib.Io
}

func NewOpener() *Opener {
	return &Opener{
		io: lib.NewIo(),
	}
}

func (o *Opener) SetIo(io *lib.Io) {
	o.io = io
}

func (o *Opener) Open(path string, environment string, siteName string, isAdmin bool) error {
	playbook := lib.NewAdHocPlaybook(map[string]string{
		"cyberduck_open.yml":    strings.TrimSpace(cyberduckOpenYml) + "\n",
		"cyberduck_bookmark.j2": strings.TrimSpace(cyberduckBookmarkJ2) + "\n",
	}, path, o.io)

	playbookArgs := []string{
		"-e", "dest=" + fmt.Sprintf("%s/cyberduck-%d.duck", path, time.Now().UnixNano()),
		"-e", "env=" + environment,
		"-e", "site=" + siteName,
		"-e", "is_admin=" + strconv.FormatBool(isAdmin),
	}

	return playbook.Run("cyberduck_open.yml", playbookArgs)
}
