package configs

import (
	"fmt"

	ini "github.com/nanitor/goini"
)

type MainConfig struct {
	SecretKey string
	Database  *DatabaseConfig
}

func LoadMainConfig(dict ini.Dict) (*MainConfig, error) {
	var err error

	ret := &MainConfig{}

	ret.SecretKey = dict.GetStringDef("main", "secret_key", "")
	if ret.SecretKey == "" {
		return nil, fmt.Errorf("secret key is empty")
	}

	ret.Database, err = DatabaseConfigsFromDict(dict)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
