package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/rest-demo/pkg/forms"
)

func CustomErrorHandler(err error, c echo.Context) {
	respCode := 500
	resp := forms.BasicResponse{}
	resp.Success = false
	resp.Message = ""

	if err == echo.ErrUnauthorized {
		respCode = http.StatusUnauthorized
	}

	if !c.Response().Committed {

		if c.Request().Method == http.MethodHead {
			err = c.NoContent(respCode)
		} else {
			resp.Message = err.Error()
			err = c.JSON(respCode, resp)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}

}
