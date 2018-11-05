package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	// defer db.Close()

	accountInitialized := db.HasTable(&Account{})
	questionTypeInitialized := db.HasTable(&QuestionType{})

	db.Debug().AutoMigrate(&Account{},
		&Questionnaire{},
		&Question{},
		&QuestionType{})

	if !questionTypeInitialized {
		db.Create(&QuestionType{Expression: "text"})
		db.Create(&QuestionType{Expression: "select"})
		db.Create(&QuestionType{Expression: "multiselect"})
		db.Create(&QuestionType{Expression: "yes/no"})
	}

	if !accountInitialized {
		account := &Account{Email: "admin@admin.com", Password: "pass"}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
		account.Password = string(hashedPassword)

		db.Create(account)
	}
}

//GetDB wrapper
func GetDB() *gorm.DB {
	return db
}
