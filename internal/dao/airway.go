package dao

import (
	"cfptoroute/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Airway struct {
	DB *gorm.DB
}

func GetAirway(db *gorm.DB) Airway {
	return Airway{db}
}

func (a Airway) GetAirwayByName(name string) ([]model.Airway, error) {
	var aw []model.Airway
	err := errors.WithStack(a.DB.Where("name = ?", name).Find(&aw).Error)
	if err != nil {
		return aw, err
	}
	return aw, nil
}

func (a Airway) GetAirwayByWhereToWhere(name string, point1 string, point2 string) ([]model.Airway, error) {
	var aw model.Airway
	err := errors.WithStack(a.DB.Where("name = ? AND point = ?", name, point1).First(&aw).Error)
	if err != nil {
		return nil, err
	}
	var aw2 model.Airway
	err = errors.WithStack(a.DB.Where("name = ? AND point = ?", name, point2).First(&aw2).Error)
	if err != nil {
		return nil, err
	}
	var aw3 []model.Airway
	if aw2.LegId > aw.LegId {
		// aw2 is after aw
		// 获取aw到aw2之间的航路
		err = errors.WithStack(a.DB.Where("name = ? AND leg_id >= ? AND leg_id <= ?", name, aw.LegId, aw2.LegId).Find(&aw3).Error)
	} else if aw2.LegId < aw.LegId {
		// aw2 is before aw
		// 获取aw2到aw之间的航路
		err = errors.WithStack(a.DB.Where("name = ? AND leg_id >= ? AND leg_id <= ?", name, aw2.LegId, aw.LegId).Order("leg_id DESC").Find(&aw3).Error)
	} else {
		aw3 = append(aw3, aw)
		aw3 = append(aw3, aw2)
	}
	if err != nil {
		return aw3, err
	}
	return aw3, nil
}
