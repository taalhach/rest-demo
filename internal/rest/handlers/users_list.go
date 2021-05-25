package handlers

import (
	"net/http"
	"time"

	"github.com/taalhach/rest-demo/internal/rest/common"

	"github.com/labstack/echo/v4"
)

const query = "SELECT id, email, role, created_at FROM USERS;"

type UserListItem struct {
	Id        int64     `json:"id" xorm:"id"`
	Email     string    `json:"email" xorm:"email"`
	Role      string    `json:"role" xorm:"role"`
	CreatedAt time.Time `json:"created_at" xorm:"created_at"`
}
type UserListResponse struct {
	Users []*UserListItem
}

func UsersList(c echo.Context) error {
	cc := c.(*common.CustomContext)

	dbSession := cc.DbSession()

	var users []*UserListItem
	err := dbSession.SQL(query).Find(&users)
	common.CheckErr(err)

	ret := UserListResponse{
		Users: users,
	}

	return c.JSON(http.StatusOK, ret)

}
