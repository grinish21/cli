package mon

import (
	"objectapi/pkg/log"
	"objectapi/pkg/mon"
	"path"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		addr   string
		script string
	}
	var options = &ClientOptions{}
	var cmd = &cobra.Command{
		Use:   "client",
		Short: "API monitor client",
		Long:  `The monitor client allows you to send api events the monitor server for testing purposes.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			options.script = args[0]
			log.Debug("run script ", options.script)
			switch path.Ext(options.script) {
			case ".json":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.addr)
				go mon.ReadJsonEvents(options.script, emitter)
				sender.SendEvents(emitter)
			case ".js":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.addr)
				vm := mon.NewEventScript(emitter)
				vm.RunScript(options.script)
				sender.SendEvents(emitter)
			case ".csv":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.addr)
				go mon.ReadCsvEvents(options.script, emitter)
				sender.SendEvents(emitter)
			default:
				log.Error("unknown file type: ", options.script)
			}
		},
	}
	cmd.Flags().StringVarP(&options.addr, "addr", "", "localhost:8080", "monitor server address")

	return cmd
}
