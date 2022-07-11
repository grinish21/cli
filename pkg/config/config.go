package config

import (
	"apigear/pkg/log"

	"github.com/spf13/viper"
)

func ReadRecentProjects() []string {
	return viper.GetStringSlice("recent")
}

func AppendRecentProject(file string) {
	// check if file is already in recent
	recent := ReadRecentProjects()
	for _, f := range recent {
		if f == file {
			log.Debugf("File %s is already in recent", file)
			return
		}
	}
	viper.Set("recent", append(recent, file))
	viper.WriteConfig()
}

func RemoveRecentFile(d string) {
	recent := ReadRecentProjects()
	for i, f := range recent {
		if f == d {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	viper.Set("recent", recent)
	viper.WriteConfig()
}