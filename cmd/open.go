package cmd

import (
	"github.com/spf13/cobra"

	"github.com/itinerisltd/trellis-cyberduck/cyberduck"
	"github.com/itinerisltd/trellis-cyberduck/lib"
)

func init() {
	opener := cyberduck.NewOpener()
	user := ""
	directory := ""

	// openCmd represents the open command
	openCmd := &cobra.Command{
		Use: "open <environment> [<site>]",
		Example: `  $ trellis-cyberduck open production example.com
  $ trellis-cyberduck open staging my-site --user admin
  $ trellis-cyberduck open staging my-site --directory project_uploads_path`,
		Short: "Open SFTP connections to Trellis servers",
		Args:  cobra.RangeArgs(1, 2),
		PreRun: func(cmd *cobra.Command, args []string) {
			io := lib.NewIoFromCobraCommand(cmd)
			opener.SetIo(io)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			maybeSite := ""
			if len(args) >= 2 {
				maybeSite = args[1]
			}

			trellis := lib.Trellis{}
			_, path, env, site, err := trellis.DetectPathAndEnvAndSite(args[0], maybeSite)
			if err != nil {
				return err
			}

			return opener.Open(path, env, site, user, directory)
		},
	}

	openCmd.Flags().StringVarP(&user, "user", "u", "web", "User to connect. Options: web|admin")
	openCmd.Flags().StringVarP(&directory, "directory", "d", "project_root", "Directory to open. Options: project_root|project_source_path|project_uploads_path|project_current_symlink_path")

	rootCmd.AddCommand(openCmd)
}
