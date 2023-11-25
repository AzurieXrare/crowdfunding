package main

import (
	"crowdfunding/auth"
	"crowdfunding/handler"
	"crowdfunding/user"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// https://github.com/go-sql-driver/mysql
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// fmt.Println(authService.GenerateToken(1001))

	// Untuk test service
	// input := user.LoginInput{
	// 	Email:    "vije@gmail.com",
	// 	Password: "testes",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// buat cek
	// userByEmail, err := userRepository.FindByEmail("vijee@gmail.com")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0 {
	// 	fmt.Println("user not found")
	// } else {
	// 	fmt.Println(userByEmail.Name)

	// }

	// buat cek masuk ato ngga secara paksa
	// userService.SaveAvatar(1, "images/1-profile.png")

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// fmt.Println("Connection to database is good")

	// var users []user.User
	// // length := len(users)
	// // fmt.Println(length)

	// db.Find(&users)
	// // length = len(users)
	// // fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=======")
	// }

	// router := gin.Default()
	// router.GET("/", handler)
	// router.Run()

}

// *gin.context merupakan handler gin
// func handler(c *gin.Context) {

// }

// user melakukan input
// handler menangkap inputan user dan dimapping kedalam struct
// service  melakukan mapping ke struct user agar dapat disimpan ke repo
// repository melakukan save struct user ke db
// db save ke db
