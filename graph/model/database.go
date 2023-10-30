package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	Dbc *gorm.DB
}

func Open() (Conn, error) {
	dsn := "host=localhost user=postgres password=admin dbname=postgres1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Conn{}, err
	}

	return Conn{Dbc: db}, nil
}
func AutoMigrate(s *Conn) error {
	//if s.db.Migrator().HasTable(&User{}) {
	//	return nil
	//}

	err := s.Dbc.Migrator().DropTable(&User{}, &Company{}, &Job{})
	if err != nil {
		return err
	}
	err = s.Dbc.Migrator().AutoMigrate(&User{}, &Company{}, &Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	return nil
}
