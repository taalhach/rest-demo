package rest

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/taalhach/rest-demo/internal/rest/common"
	"github.com/taalhach/rest-demo/internal/rest/handlers"
	"github.com/taalhach/rest-demo/internal/rest/models"
	"github.com/taalhach/rest-demo/pkg/forms"
)

const port = 8081

var serveApiCmd = &cobra.Command{
	Use:   "serve_api",
	Short: "servers api",
	Long:  fmt.Sprintf("servers rest ums api on localhost port %v", port),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := loadConfigs()
		if err != nil {
			fmt.Printf("failed to load configs: %v", err)
			os.Exit(1)
		}

		e := echo.New()

		e.HTTPErrorHandler = func(err error, c echo.Context) {
			respCode := 500
			resp := forms.BasicResponse{}
			resp.Success = false
			resp.Message = ""
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

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("method=${method} uri=${uri} status=${status} time=${latency_human}\n"),
		}))

		// middleware to recover from panics calls HttpErrorHandler
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize:         1 << 10, // 1 KB
			DisableStackAll:   true,
			DisablePrintStack: true,
		}))

		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
		e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc := &common.CustomContext{
					Context: c,
					Engine:  Engine,
					MainCfg: MainConfigs,
				}

				//	db cleanup before writing response
				cc.Response().Before(func() {
					cc.Cleanup()
				})

				if strings.Contains(cc.Request().Header.Get("Content-Encoding"), "gzip") {
					// Decompress the stream
					cc.Request().Body, err = gzip.NewReader(cc.Request().Body)
					if err != nil {
						return err
					}

					defer cc.Request().Body.Close()
				}

				return h(cc)
			}
		})
		e.Use(middleware.JWTWithConfig(middleware.JWTConfig{

			SigningKey: []byte(MainConfigs.SecretKey),
			Skipper: func(c echo.Context) bool {
				if strings.HasSuffix(c.Path(), "/register") {
					return true
				}

				if strings.HasSuffix(c.Path(), "/login") {
					return true
				}
				return false
			},
			SuccessHandler: func(c echo.Context) {

				token, ok := c.Get("user").(*jwt.Token)
				if !ok {
					return
				}

				claims := token.Claims.(jwt.MapClaims)
				uuid := claims["sub"]
				if uuid == "" {
					return
				}

				cc := c.(*common.CustomContext)
				dbSession := cc.DbSession()

				var user models.User
				has, err := dbSession.Where("uuid = ?", uuid).Get(&user)
				if !has || err != nil {
					return
				}

				cc.User = &user
			},
		}))

		e.POST("/login", handlers.Login)
		e.POST("/register", handlers.Register)
		e.GET("/users", handlers.UsersList)

		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	},
}

func init() {
	rootCommand.AddCommand(serveApiCmd)
}