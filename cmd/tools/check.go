package tools

import (
	"apigear/pkg/log"
	"apigear/pkg/spec"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewCheckCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "check",
		Short: "check document",
		Long:  `check documents and report errors`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			var file = args[0]
			switch filepath.Ext(file) {
			case ".json", ".yaml":
				result, err := spec.CheckFile(file)
				if err != nil {
					log.Warnf("failed to check json file %s: %s", file, err)
					break
				}
				if result.Valid() {
					fmt.Printf("valid: %s\n", file)
				} else {
					fmt.Printf("invalid: %s\n", file)
					for _, desc := range result.Errors() {
						fmt.Println(desc.String())
					}
				}
			case ".csv":
				err := spec.CheckCsvFile(file)
				if err != nil {
					log.Warnf("failed to check csv file %s: %s", file, err)
				} else {
					fmt.Printf("valid: %s\n", file)
				}
			case ".ndjson":
				err := spec.CheckNdjsonFile(file)
				if err != nil {
					log.Warnf("failed to check ndjson file %s: %s", file, err)
				} else {
					fmt.Printf("valid: %s\n", file)
				}
			case ".idl":
				err := spec.CheckIdlFile(file)
				if err != nil {
					log.Warnf("failed to check idl file %s: %s", file, err)
				} else {
					fmt.Printf("valid: %s\n", file)
				}
			default:
				fmt.Printf("unknown file type %s", file)
			}
		},
	}
	return cmd
}