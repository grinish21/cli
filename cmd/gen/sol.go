package gen

import (
	"fmt"
	"sync"

	"github.com/apigear-io/lib/log"
	"github.com/apigear-io/lib/sol"

	"github.com/spf13/cobra"
)

func NewSolutionCommand() *cobra.Command {
	var file string
	var watch bool
	var cmd = &cobra.Command{
		Use:     "solution [solution-file]",
		Short:   "generate SDK using a solution document",
		Aliases: []string{"sol", "s"},
		Args:    cobra.ExactArgs(1),
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.OnReport(func(report *log.ReportEntry) {
				fmt.Print(report.Message)
			})
			file = args[0]
			doc, err := sol.ReadSolutionDoc(file)
			if err != nil {
				return err
			}
			runner := sol.NewRunner()
			if watch {
				var wg = sync.WaitGroup{}
				wg.Add(1)
				done, err := runner.StartWatch(file, doc)
				if err != nil {
					return err
				}
				wg.Wait()
				done <- true
			} else {
				err := runner.RunDoc(file, doc)
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
