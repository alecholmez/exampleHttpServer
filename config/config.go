package config

import "github.com/spf13/viper"

// NewConfig ...
func NewConfig(dir, configName string) error {
	viper.SetConfigName(configName)
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
