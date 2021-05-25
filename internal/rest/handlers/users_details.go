package handlers

import (
	"net/http"
	"time"

	"github.com/taalhach/rest-demo/pkg/forms"

	"github.com/taalhach/rest-demo/internal/rest/common"

	"github.com/labstack/echo/v4"
)

const userDetailsQuery = "SELECT id, email, role, created_at FROM USERS WHERE id = ?"

type UserDetailsItem struct {
	Id        int64     `json:"id" xorm:"id"`
	Email     string    `json:"email" xorm:"email"`
	Role      string    `json:"role" xorm:"role"`
	CreatedAt time.Time `json:"created_at" xorm:"created_at"`
}

type UserDetailsResponse struct {
	forms.BasicResponse
	User *UserDetailsItem
}

func UsersDetails(c echo.Context) error {
	cc := c.(*common.CustomContext)

	id := c.Param("user-id")
	dbSession := cc.DbSession()

	var user UserDetailsItem
	has, err := dbSession.SQL(userDetailsQuery, id).Get(&user)
	common.CheckErr(err)

	var ret UserDetailsResponse
	if !has {
		ret.Success = false
		ret.Message = "user not found"
		return c.JSON(http.StatusNotFound, ret)
	}

	ret.Success = true
	ret.User = &user

	return c.JSON(http.StatusOK, ret)
}
