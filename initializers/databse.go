package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("unable to Load DB credentials")
	}

	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		myEnv["host"],
		myEnv["user"],
		myEnv["password"],
		myEnv["dbname"],
		myEnv["port"],
		myEnv["sslmode"],
		myEnv["TimeZone"])

	dsn := dbString
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection failed")
	}
	if err == nil {
		fmt.Println("connected to DB ")
	}

}
