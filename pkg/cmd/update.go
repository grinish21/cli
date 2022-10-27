package cmd

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/up"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var force bool
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update the program",
		Long:  `check and update the program to the latest version`,
		Run: func(cmd *cobra.Command, args []string) {
			repo := "apigear-io/cli-releases"
			version := config.BuildVersion()
			u, err := up.NewUpdater(repo, version)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			release, err := u.Check()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			if release == nil {
				cmd.Println("no new release available")
				return
			}
			cmd.Printf("New release %s available.\n", release.Version())
			cmd.Printf("See %s.\n", release.URL)
			if !force {
				result, err := pterm.DefaultInteractiveConfirm.Show("do you want to update?")
				if err != nil {
					cmd.PrintErrln(err)
					return
				}
				if !result {
					return
				}
			}
			fmt.Printf("updating to %s\n", release.Version())
			err = u.Update(release)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "force update")
	return cmd
}
