package rest

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	ini "github.com/nanitor/goini"
	"github.com/taalhach/rest-demo/internal/rest/configs"
)

const (
	envKey = "REST_TASK_SETTINGS"
)

var (
	MainConfigs *configs.MainConfig
	Engine      *xorm.Engine
)

func loadConfigs() (*configs.MainConfig, error) {

	file := os.Getenv(envKey)
	if file == "" {
		fmt.Printf("Missing env variable: %v", envKey)
		os.Exit(1)
	}

	dict, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	MainConfigs, err = configs.LoadMainConfig(dict)
	if err != nil {
		return nil, err
	}

	// make connection
	Engine = MainConfigs.Database.MustGetEngine()

	return MainConfigs, err
}
