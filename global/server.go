package global

import (
	"cfptoroute/handler"
	"math/rand"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {

				if err, ok := r.(error); ok {
					c.Error(errors.WithStack(err))
				} else if rr, ok := r.(interface{}); ok {
					c.Error(errors.Errorf("%v", rr))
				}
			}
		}()
		return next(c)
	}
}

func CreateWebServer() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, context echo.Context) {
		id := RandSeq(10)
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				_ = handler.ErrorResponse(context, handler.StatusResourceNotFound)
			} else {
				_ = context.JSON(he.Code, echo.Map{
					"code":    he.Code,
					"message": he.Message,
				})
			}
		} else if err != nil {
			_ = context.JSON(500, echo.Map{
				"code":     -1100,
				"message":  "Unexceptional error",
				"trace_id": id,
			})
			WebLog.WithField("trace_id", id).Errorf("%+v", err)
		}
	}

	e.Use(Recover)
	return e
}
