package main

import (
	"cfptoroute/global"
	"cfptoroute/internal/controller"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dbInit(dbPath string) error {
	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}
	global.DB = DB
	return nil
}

func main() {
	dbInit(global.DbPath)
	global.LogInit()
	httpServer := global.CreateWebServer()
	controller.RouterInit(httpServer)
	httpServer.Start(":8080")
	// res, _ := service.SegmentToPointsList(service.ParseCFPRoute("APESO W615 IDAXI R473 BEMAG V5 CON A599 POU"))
	// fmt.Println(res)
	// for _, v := range res {
	// 	fmt.Println(v.Name)
	// }
	defer func() {
		if err := recover(); err != nil {
			global.CoreLog.Errorf("%+v", err)
		}
	}()
}
