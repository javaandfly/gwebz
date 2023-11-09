package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// read configuration file
func ReadConfig(filePath string, config interface{}) error {
	vp := viper.New()
	// set path
	vp.SetConfigFile(filePath)

	// compatible with toml type because .conf and .cxt cannot be parsed
	fileType := filepath.Ext(filePath)
	if fileType == ".conf" || fileType == ".cfg" {
		vp.SetConfigType("toml")
	}

	vp.AutomaticEnv()

	//the . in the field is replaced by _
	replacer := strings.NewReplacer(".", "_")
	vp.SetEnvKeyReplacer(replacer)

	// read file
	if err := vp.ReadInConfig(); err != nil {
		err = fmt.Errorf("read config file failed %v", err)
		return err
	}

	//replace according to the given structure
	if err := vp.Unmarshal(config); err != nil {
		err = fmt.Errorf("unmarshal config file failed %v", err)
		return err
	}
	return nil

}
