package cmd

import (
	"github.com/spf13/cobra"

	"github.com/itinerisltd/trellis-cyberduck/cyberduck"
	"github.com/itinerisltd/trellis-cyberduck/lib"
)

func init() {
	opener := cyberduck.NewOpener()
	user := ""

	// openCmd represents the open command
	openCmd := &cobra.Command{
		Use: "open <environment> [<site>]",
		Example: `  $ trellis-cyberduck open production example.com
  $ trellis-cyberduck open staging my-site --admin
`,
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

			return opener.Open(path, env, site, user)
		},
	}

	openCmd.Flags().StringVarP(&user, "user", "u", "web", "Connect as web or admin user. Option: web|admin")

	rootCmd.AddCommand(openCmd)
}
