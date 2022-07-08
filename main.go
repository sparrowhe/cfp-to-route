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

// func CopyDatabaseToMem(dbPath string) error {
// 	// 将数据库复制到内存中
// 	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}
// 	DBMem, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 	DBMem.AutoMigrate(&model.Airway{}, &model.Airport{})
// 	// 从磁盘写入数据
// 	var waypoints []model.Airway
// 	DB.Table("airways").Find(&waypoints)
// 	// 分批写入内存
// 	for i := 0; i < len(waypoints); i += 900 {
// 		DBMem.Create(waypoints[i : i+900])
// 	}
// 	global.DB = DBMem
// 	return nil
// }

func main() {
	dbInit(global.DbPath)
	global.LogInit()
	httpServer := global.CreateWebServer()
	controller.RouterInit(httpServer)
	httpServer.Start(":8080")
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
