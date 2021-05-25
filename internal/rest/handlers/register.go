package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/rest-demo/internal/rest/common"
	"github.com/taalhach/rest-demo/internal/rest/models"
	"github.com/taalhach/rest-demo/pkg/forms"
	"github.com/taalhach/rest-demo/pkg/items"
)

type RegisterForm struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func Register(c echo.Context) error {
	cc := c.(*common.CustomContext)

	form := new(RegisterForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if !strings.EqualFold(form.Password, form.ConfirmPassword) {
		ret := forms.BasicResponse{
			Success: false,
			Message: "both password should match",
		}

		return c.JSON(http.StatusOK, ret)
	}

	dbSession := cc.DbSession()
	utcNow := time.Now().UTC()

	var user models.User
	has, err := dbSession.Where("email = ?", form.Email).Get(&user)
	common.CheckErr(err)

	if has {
		ret := forms.BasicResponse{
			Success: false,
			Message: "User already exists",
			Errors:  map[string]string{"email": "user already exists"},
		}

		return c.JSON(http.StatusOK, ret)
	}

	encryptedPass, err := common.Encrypt(cc.MainCfg.SecretKey, form.Password)
	common.CheckErr(err)

	uuid, err := common.NewGuid()
	common.CheckErr(err)

	user.Password = encryptedPass
	user.Email = form.Email
	user.CreatedAt = utcNow
	user.LastUpdatedAt = utcNow
	user.Uuid = uuid
	user.Role = items.UserRole

	_, err = dbSession.Insert(&user)
	common.CheckErr(err)

	err = dbSession.Commit()
	common.CheckErr(err)

	ret := forms.BasicResponse{
		Success: true,
		Message: "user registered successfully",
	}

	return c.JSON(http.StatusOK, ret)
}
