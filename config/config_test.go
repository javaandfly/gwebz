package config

import (
	"testing"
)

func TestReadConfig(t *testing.T) {

	tests := &struct {
		Name  string `mapstructure:"name"`
		Param struct {
			Age int64 `mapstructure:"age"`
		} `mapstructure:"param"`
	}{
		Name: "xiaodong",
		Param: struct {
			Age int64 `mapstructure:"age"`
		}{
			Age: 18,
		},
	}

	cfgPojo := &struct {
		Name  string `mapstructure:"name"`
		Param struct {
			Age int64 `mapstructure:"age"`
		} `mapstructure:"param"`
	}{}
	// c := &config.Config{}
	err := ReadConfig("/Users/xiaodong/go/src/myproject/go-web-formwork/config/config_test.yaml", cfgPojo)
	if err != nil {
		panic(err)
	}
	if *cfgPojo != *tests {
		t.Errorf("ReadConfig() = %v, tests %v", cfgPojo, tests)
	}
}
