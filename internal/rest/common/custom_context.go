package common

import (
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo/v4"
	"github.com/taalhach/rest-demo/internal/rest/configs"
	"github.com/taalhach/rest-demo/internal/rest/models"
)

type CustomContext struct {
	echo.Context
	MainCfg   *configs.MainConfig
	Engine    *xorm.Engine
	dbSession *xorm.Session
	User      *models.User
}

func (cc *CustomContext) DbSession() *xorm.Session {
	if cc.dbSession == nil {
		cc.dbSession = NewDbSession(cc.Engine)
	}

	return cc.dbSession
}

func NewDbSession(engine *xorm.Engine) *xorm.Session {
	session := engine.NewSession()
	err := session.Begin()
	if err != nil {
		panic(err)
	}
	return session
}

func (this *CustomContext) Cleanup() {
	if this.dbSession != nil {
		this.dbSession.Close()
		this.dbSession = nil
	}
}
