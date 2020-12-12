package db

import (
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model

		Email    string `gorm:"type:varchar(100)"`
		Password string `gorm:"type:varchar(200)"`
	}

	PreferenceCategory struct {
		gorm.Model

		CategoryName        string `gorm:"type:varchar(300)"`
		CategoryHelpText    string `gorm:"type:varchar(512)"`
		CategoryHelpMessage string `gorm:"type:varchar(512)"`
		DisplayOrder        uint
		Enabled             bool
		Lat                 string
		Long                string
		CityName            string
	}

	PreferenceCategoryOption struct {
		ID             uint   `gorm:"primary_key"`
		OptionName     string `gorm:"type:varchar(100)"`
		OptionText     string `gorm:"type:varchar(200)"`
		Enabled        bool   `gorm:"default:true"`
		PrefCategoryID uint
	}

	UserPreferences struct {
		ID                   uint        `gorm:"primary_key"`
		UserID               interface{} `gorm:"type:integer"`
		PrefCategoryOptionID uint
	}
)
