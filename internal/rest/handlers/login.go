package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/taalhach/rest-demo/internal/rest/common"
	"github.com/taalhach/rest-demo/internal/rest/models"
	"github.com/taalhach/rest-demo/pkg/forms"
)

type UserSignInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignInResponse struct {
	forms.BasicResponse
	AccessToken string `json:"access_token"`
}

func Login(c echo.Context) error {
	cc := c.(*common.CustomContext)

	form := new(UserSignInForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	dbSession := cc.DbSession()

	var user models.User
	has, err := dbSession.Where("email = ?", form.Email).Get(&user)
	if err != nil {
		return err
	}

	var ret UserSignInResponse
	if !has {
		ret.Success = false
		ret.Message = "user not found"
		return c.JSON(http.StatusNotFound, ret)
	}

	matched, err := checkAuth(&user, form.Password, cc.MainCfg.SecretKey)
	if err != nil {
		return err
	}

	if !matched {
		ret.Success = false
		ret.Message = "invalid password"
		return c.JSON(http.StatusUnauthorized, ret)
	}

	expiry := time.Now().UTC().Add(72 * time.Hour)
	accessToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = expiry.Unix()
	claims["sub"] = user.Uuid
	tokenStr, err := accessToken.SignedString([]byte(cc.MainCfg.SecretKey))
	common.CheckErr(err)

	ret.Success = true
	ret.Message = "successful sign in"
	ret.AccessToken = tokenStr

	return c.JSON(http.StatusOK, ret)
}

func checkAuth(user *models.User, password, secret string) (bool, error) {
	pass, err := common.Decrypt(secret, user.Password)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(password, pass), nil
}
