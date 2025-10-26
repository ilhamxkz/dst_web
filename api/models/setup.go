package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// 192.168.100.135 (waditra)
	// 192.168.29.135 (kost)
	// 192.168.1.4 (rumah ilham)
	// 192.168.137.139 (hapitt)
	// 192.168.70.17 (ethernet sekolah)
	// 192.168.18.53 (rumah daris)
	// 127.0.0.1 (localhost)

	dsn := "root:@tcp(127.0.0.1:3306)/webportal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	DB = db

	fmt.Println("Database connected")
}

// Optional: Auto-migrate helper
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{}, &Role{}, &Menu{}, &Akses{}, &Order{}, &Post{}, &BlogPost{}, &BlogComment{}, &ProjectImage{}, &Pertanyaan{}, &Testimoni{},
		&Pelatihan{}, &Statistic{},
	)
}
