package cmd

import (
	"fmt"
	"github.com/itinerisltd/trellis-cyberduck/lib"
	"github.com/mitchellh/cli"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
	trellisLib "trellis-cli/trellis"
)

const cyberduckOpenYml = `
---
- name: 'Trellis CLI: Dump database credentials'
  hosts: web:&{{ env }}
  remote_user: "{{ web_user }}"
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
	<string>{{ web_user }}</string>
	<key>Path</key>
	<string>{{ project_root | default(www_root + '/' + item.key) | regex_replace('^~\/','') }}/{{ item.current_path | default('current') }}</string>
</dict>
</plist>
`

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open <environment> [<site>]",
	Example: `  $ trellis-cyberduck open production example.com
  $ trellis-cyberduck open staging my-site
`,
	Short: "Open SFTP connections to Trellis servers",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure we are inside trellis root and activate virtualenv.
		project := &trellisLib.Project{}
		trellis := trellisLib.NewTrellis(project)
		if err := trellis.LoadProject(); err != nil {
			return err
		}

		// Validate environment exist.
		environment := args[0]
		if err := trellis.ValidateEnvironment(environment); err != nil {
			return err
		}

		// Validate or detect site.
		siteNameArg := ""
		if len(args) >= 2 {
			siteNameArg = args[1]
		}
		siteName, siteNameErr := trellis.FindSiteNameFromEnvironment(environment, siteNameArg)
		if siteNameErr != nil {
			return siteNameErr
		}

		// Open!
		ui := &cli.ColoredUi{
			ErrorColor: cli.UiColorRed,
			Ui: &cli.BasicUi{
				Reader:      os.Stdin,
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
			},
		}

		playbook := lib.AdHocPlaybook{
			Files: map[string]string{
				"cyberduck_open.yml":    cyberduckOpenYml,
				"cyberduck_bookmark.j2": strings.TrimSpace(cyberduckBookmarkJ2) + "\n",
			},
			Playbook: lib.Playbook{
				Root: trellis.Path,
				UI:   ui,
			},
		}

		playbookArgs := []string{
			"-e", "env=" + environment,
			"-e", "site=" + siteName,
			"-e", "dest=" + fmt.Sprintf("%s/cyberduck-%d.duck", trellis.Path, time.Now().UnixNano()),
		}

		return playbook.Run("cyberduck_open.yml", playbookArgs)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
