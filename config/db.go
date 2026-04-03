package config

import (
    "fmt"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
    godotenv.Load()
    
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
    
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Gagal koneksi ke database!")
    }

    DB = database
    fmt.Println("Database Connected!")
}