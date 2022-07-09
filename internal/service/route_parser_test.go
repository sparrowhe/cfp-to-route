package service

import (
	"cfptoroute/global"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Benchmark_ParseCFPRoute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCFPRoute("ZBAA BOTPU W47 LOVRA W540 DOVOP H138 ONEBA G212 JTG W236 CTU B213 LXA ZULS")
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
	Init("../../../../navdata.db")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SegmentToPointsList(ParseCFPRoute("ZBAA BOTPU W47 LOVRA W540 DOVOP H138 ONEBA G212 JTG W236 CTU B213 LXA ZULS"))
	}
}

func Benchmark_MemoryConvert(b *testing.B) {
	Init("../../../../navdata.db")
	wp, _ := SegmentToPointsList(ParseCFPRoute("ZBAA BOTPU W47 LOVRA W540 DOVOP H138 ONEBA G212 JTG W236 CTU B213 LXA ZULS"))
	SaveCache("ZBAA BOTPU W47 LOVRA W540 DOVOP H138 ONEBA G212 JTG W236 CTU B213 LXA ZULS", wp)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoadCache("ZBAA BOTPU W47 LOVRA W540 DOVOP H138 ONEBA G212 JTG W236 CTU B213 LXA ZULS")
	}
}
