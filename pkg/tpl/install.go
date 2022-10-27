package tpl

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

// InstallTemplate installs template template registry into the cache
func InstallTemplate(name string) error {
	// check if name is a local dst in the cache
	dst := helper.Join(config.TemplateCacheDir(), name)
	if helper.IsDir(dst) {
		return fmt.Errorf("template %s already exists", name)
	}
	// name should exists in the registry
	reg, err := ReadRegistry()
	if err != nil {
		return err
	}
	for _, t := range reg.Entries {
		if t.Name == name {
			// clone the repo
			err := git.Clone(t.Git, dst)
			if err != nil {
				return err
			}
			if t.Latest != "" {
				err = git.CheckoutCommit(dst, t.Latest)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return fmt.Errorf("template %s not found in registry", name)
}
