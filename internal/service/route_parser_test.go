package service

import (
	"cfptoroute/global"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Benchmark_ParseCFPRoute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCFPRoute("APESO W615 IDAXI R473 BEMAG V5 CON A599 POU")
	}
}

func Init(dbPath string) error {
	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}
	global.DB = DB
	return nil
}

func Benchmark_SegmentToPointsList(b *testing.B) {
	Init("../../navdata.db")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SegmentToPointsList(ParseCFPRoute("BOTPU W47 LOVRA W540 DOVOP H138 ONEBA G212 JTG W236 CTU B213 LXA"))
	}
}
