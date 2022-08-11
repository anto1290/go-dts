package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	dsn := "host=ec2-34-235-31-124.compute-1.amazonaws.com user=mfklocoprpptvk password=c81c464c714daa0be24653fc2dfab0e71af052d3a209c0cd63899dccfb6edd09 dbname=d4187e3mcn1mk2 port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&List{})
	DB = database
}
