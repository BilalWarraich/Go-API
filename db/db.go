package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ReaderDB struct {
	gormDB *gorm.DB
}

type WriterDB struct {
	*ReaderDB
}

var (
	reader *ReaderDB
	writer *WriterDB
)

func Reader(connectionStr string) (*ReaderDB, error) {
	if reader != nil {
		return reader, nil
	}

	gormDB, err := new(connectionStr)
	if err != nil {
		return nil, err
	}

	return &ReaderDB{gormDB: gormDB}, nil
}

func Writer(connectionStr string) (*WriterDB, error) {
	if writer != nil {
		return writer, nil
	}

	gormDB, err := new(connectionStr)
	if err != nil {
		return nil, err
	}

	// insert dummy user for testing purpose
	var user User
	if notFound := gormDB.First(&user, 1).RecordNotFound(); notFound {
		if err = gormDB.Exec("insert into users(id,email,password) values(?,?,?)", 1, "clykk@gmail.com", "frdeswaq").Error; err != nil {
			return nil, err
		}
	}

	return &WriterDB{&ReaderDB{gormDB: gormDB}}, nil
}

func new(connectionStr string) (*gorm.DB, error) {
	gormDB, err := gorm.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}
	gormDB.DB().SetMaxOpenConns(90)
	gormDB.DB().SetMaxIdleConns(0)

	gormDB.AutoMigrate(
		&User{},
		&PreferenceCategory{},
		&PreferenceCategoryOption{},
		&UserPreferences{},
	)

	// disable Logger
	gormDB.LogMode(false)

	return gormDB, nil
}
