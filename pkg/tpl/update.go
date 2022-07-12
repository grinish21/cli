package tpl

import (
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
)

func UpdateTemplate(name string) error {
	// update template
	target := filepath.Join(GetPackageDir(), name)
	log.Infof("update template %s", name)
	err := ExecGit([]string{"pull"}, target)
	if err != nil {
		log.Warnf("failed to update template %s", name)

	}
	return err
}
