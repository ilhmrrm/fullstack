package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ilhmrrm/fullstack/api/controllers"
	"github.com/ilhmrrm/fullstack/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	server       = controllers.Server{}
	userInstance = models.User{}
	postInstance = models.Post{}
)

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatal(err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}

	// if TestDbDriver == "postgres" {
	// 	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
	// 	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	// 	if err != nil {
	// 		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
	// 		log.Fatal("This is the error:", err)
	// 	} else {
	// 		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	// 	}
	// }
}

func resfreshUserTables() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	resfreshUserTables()

	user := models.User{
		Nickname: "Tri",
		Email:    "tri@gmail.com",
		Password: "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	return user, nil
}

func seedUsers() error {
	users := []models.User{
		models.User{
			Nickname: "user",
			Email:    "user@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Kenny Morris",
			Email:    "kenny@gmail",
			Password: "password",
		},
	}

	// access the users
	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndPostTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {
	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}

	user := models.User{
		Nickname: "mei",
		Email:    "mei@gmail.com",
		Password: "password",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}

	post := models.Post{
		Title:    "ini title mei",
		Content:  "ini content mei",
		AuthorID: user.ID,
	}

	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func seedUserAndPosts() ([]models.User, []models.Post, error) {
	var err error

	if err != nil {
		return []models.User{}, []models.Post{}, err
	}

	users := []models.User{
		models.User{
			Nickname: "pasya",
			Email:    "pasya@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "bintang",
			Email:    "bintang@gmail.com",
			Password: "password",
		},
	}

	posts := []models.Post{
		models.Post{
			Title:   "Title pasya",
			Content: "Content pasya",
		},
		models.Post{
			Title:   "Title bintang",
			Content: "Content bintang",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}

	return users, posts, nil
}
