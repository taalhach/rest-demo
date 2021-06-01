package rest

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/rest-demo/pkg/forms"
)

var customErrHandler = func(err error, c echo.Context) {
	var resp forms.BasicResponse
	respCode := 500

	// generally internal error should not only be visible in development mode
	// but for this task environment config are skipped dump error message
	if he, ok := err.(*echo.HTTPError); ok {
		respCode = he.Code
		resp.Message = fmt.Sprintf("%v", he.Message)

		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)

			resp.Message += " - ErrorMessageInternal: " + he.Internal.Error()
		}
	} else {
		resp.Message = http.StatusText(respCode)
	}

	// 4 KB stack
	stack := make([]byte, 4<<10)
	length := runtime.Stack(stack, false)
	fmt.Printf("[RECOVER From Exception]: %v %s\n", err, stack[:length])

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(respCode)
		} else {
			err = c.JSON(respCode, resp)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}

}
