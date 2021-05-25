package handlers

import (
	"net/http"
	"time"

	"github.com/taalhach/rest-demo/internal/rest/common"

	"github.com/labstack/echo/v4"
)

type UserProfileItemResponse struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func UsersProfile(c echo.Context) error {
	cc := c.(*common.CustomContext)

	ret := UserProfileItemResponse{
		Id:        cc.User.Id,
		Email:     cc.User.Email,
		Role:      cc.User.Role,
		CreatedAt: cc.User.CreatedAt,
	}
	return c.JSON(http.StatusOK, ret)
}
