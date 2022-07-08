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
	for i := 0; i < b.N; i++ {
		SegmentToPointsList(ParseCFPRoute("SULAS V17 UBDOB V18 XEBUL G471 PLT A599 ELNEX G204 MULOV V73 SUPAR B221 NINAS G327 LAMEN A593 SADLI Y590 ELGEP Y722 CJU A586 TENAS B467 NULAR L771 DITOR T575 LEMBA P177 TK T657 ERNIK B240 ENM J179 MDO J605 BKA J195 ANN J502 YYJ"))
	}
}
