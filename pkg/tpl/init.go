package tpl

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	packageDir := path.Join(home, ".apigear", "templates")
	viper.SetDefault("packageDir", packageDir)
	os.MkdirAll(packageDir, 0755)
}

func GetPackageDir() string {
	return viper.GetString("packageDir")
}
