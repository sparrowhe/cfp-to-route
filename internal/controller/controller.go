package controller

import (
	"cfptoroute/global"
	"cfptoroute/handler"
	"cfptoroute/internal/dao"
	"cfptoroute/internal/service"

	"github.com/labstack/echo/v4"
)

func RouterInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.GET("/convert", func(c echo.Context) error {
		token := c.QueryParam("token")
		route := c.QueryParam("route")
		h := handler.NewResponser(c)
		if token != global.Token {
			return h.Error(handler.StatusPermissionDenied)
		}
		if route == "" {
			return h.Error(handler.StatusInvalidData)
		}
		data, err := service.SegmentToPointsList(service.ParseCFPRoute(route))
		if err != nil {
			return err
		}
		return h.Data(data)
	})
	e.GET("/airportInfo", func(c echo.Context) error {
		token := c.QueryParam("token")
		airport := c.QueryParam("airport")
		AirportDao := dao.GetAirport(global.DB)
		h := handler.NewResponser(c)
		if token != global.Token {
			return h.Error(handler.StatusPermissionDenied)
		}
		if airport == "" {
			return h.Error(handler.StatusInvalidData)
		}
		data, err := AirportDao.GetAirportByIcao(airport)
		if dao.NotFound(err) {
			return h.Error(handler.StatusResourceNotFound)
		} else if err != nil {
			return err
		}
		return h.Data(data)
	})
}
