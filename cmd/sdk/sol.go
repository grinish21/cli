package sdk

import (
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sol"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

type SolutionOptions struct {
	file  string
	watch bool
}

func runSolution(file string) error {
	log.Infof("run solution %s", file)
	doc, err := sol.ReadSolutionDoc(file)
	if err != nil {
		log.Errorf("error reading solution: %s", err)
		return err
	}
	rootDir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		return err
	}
	proc := sol.NewSolutionRunner(rootDir, doc)
	proc.Run()
	return nil
}

func watchSol(options *SolutionOptions) {
	err := runSolution(options.file)
	if err != nil {
		log.Fatal(err)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Infof("file %s modified", event.Name)
					err := runSolution(options.file)
					if err != nil {
						log.Errorf("failed to run solution: %s", err)
					}
				}
			case err := <-watcher.Errors:
				log.Error(err)
			}
		}
	}()
	err = watcher.Add(options.file)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func NewSolutionCommand() *cobra.Command {
	var options = &SolutionOptions{}
	var cmd = &cobra.Command{
		Use:     "solution [file to run]",
		Short:   "generate code using a solution",
		Aliases: []string{"sol", "s"},
		Args:    cobra.ExactArgs(1),
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			options.file = args[0]
			if options.watch {
				watchSol(options)
			} else {
				err := runSolution(options.file)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&options.watch, "watch", "", false, "watch solution file for changes")
	return cmd
}
