package configs

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	ini "github.com/nanitor/goini"
)

type DatabaseConfig struct {
	Name      string
	Host      string
	Port      int
	Password  string
	User      string
	ShowError bool
	ShowSql   bool
	ShowDebug bool
}

func DatabaseConfigsFromDict(dict ini.Dict) (*DatabaseConfig, error) {
	section := "database"
	configs := &DatabaseConfig{}
	configs.Host = dict.GetStringDef(section, "host", "")
	if configs.Host == "" {
		configs.Host = "localhost"
	}

	configs.Name = dict.GetStringDef(section, "name", "")
	if configs.Name == "" {
		return nil, fmt.Errorf("missing database name")
	}

	configs.Port = dict.GetIntDef(section, "port", 5432)

	configs.Password = dict.GetStringDef(section, "password", "")
	configs.User = dict.GetStringDef(section, "user", "")
	return configs, nil
}

func (c *DatabaseConfig) ConnString() string {
	options := []string{}

	if len(c.Host) > 0 {
		options = append(options, fmt.Sprintf("host=%s", c.Host))
	}

	if c.Port > 0 {
		options = append(options, fmt.Sprintf("port=%d", c.Port))
	}

	if len(c.Name) > 0 {
		options = append(options, fmt.Sprintf("dbname=%s", c.Name))
	}

	if len(c.User) > 0 {
		options = append(options, fmt.Sprintf("user=%s", c.User))
	}

	if len(c.Password) > 0 {
		options = append(options, fmt.Sprintf("password=%s", c.Password))
	}

	options = append(options, "sslmode=disable")

	return strings.Join(options, " ")
}

func (this *DatabaseConfig) MustGetEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("postgres", this.ConnString())
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	engine.TZLocation = time.UTC

	engine.SetMaxOpenConns(10)
	engine.SetMaxIdleConns(5)

	if this.ShowSql {
		engine.ShowSQL(true)
	}

	if this.ShowError {
		engine.Logger().SetLevel(core.LOG_ERR)
	} else if this.ShowDebug {
		engine.Logger().SetLevel(core.LOG_DEBUG)
	} else {
		engine.Logger().SetLevel(core.LOG_OFF)
	}

	return engine
}
