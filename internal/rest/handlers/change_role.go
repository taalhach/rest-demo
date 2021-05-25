package handlers

import (
	"net/http"

	"github.com/taalhach/rest-demo/pkg/forms"

	"github.com/taalhach/rest-demo/pkg/items"

	"github.com/taalhach/rest-demo/internal/rest/models"

	"github.com/taalhach/rest-demo/internal/rest/common"

	"github.com/labstack/echo/v4"
)

func ChangeRole(c echo.Context) error {
	cc := c.(*common.CustomContext)

	id := c.Param("user-id")
	dbSession := cc.DbSession()

	var user models.User
	has, err := dbSession.Where("id = ?", id).Get(&user)
	common.CheckErr(err)

	var ret forms.BasicResponse
	if !has {
		ret.Success = false
		ret.Message = "user not found"
		return c.JSON(http.StatusNotFound, ret)
	}

	user.Role = items.AdminRole

	_, err = dbSession.Where("id = ?", user.Id).Cols("role").Update(&user)
	common.CheckErr(err)

	err = dbSession.Commit()
	common.CheckErr(err)

	ret.Success = true
	ret.Message = "user role is set to admin successfully"
	return c.JSON(http.StatusOK, ret)
}
