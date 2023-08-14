package tools

import (
	"github.com/spf13/viper"
)

func ReadFromYaml(path string, target any) error {
	return readConfigFile(path, target, "yaml")
}

func readConfigFile(path string, target any, t string) error {
	app := viper.New()
	app.SetConfigFile(path)
	app.SetConfigType(t)
	if err := app.ReadInConfig(); err != nil {
		return err
	}
	return app.Unmarshal(target)
}
