package sdk

import (
	"github.com/apigear-io/cli/pkg/sol"

	"github.com/spf13/cobra"
)

func NewSolutionCommand() *cobra.Command {
	var file string
	var watch bool
	var cmd = &cobra.Command{
		Use:     "solution [file to run]",
		Short:   "generate code using a solution",
		Aliases: []string{"sol", "s"},
		Args:    cobra.ExactArgs(1),
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			file = args[0]
			if watch {
				sol.WatchSolution(file)
			} else {
				err := sol.RunSolution(file)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&watch, "watch", "", false, "watch solution file for changes")
	return cmd
}
