package main

import (
	"bwastartup/handler"
	"bwastartup/mahasiswa"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "root:12345678@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewServices(userRepository)
	userHandler := handler.NewUserHandler(userService)


	mahasiswaRepository := mahasiswa.NewRepository(db)
	mahasiswaService := mahasiswa.NewServices(mahasiswaRepository)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaService)



	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users" , userHandler.RegisterUser)
	api.POST("/mahasiswa" , mahasiswaHandler.Register)
	router.Run()
	//input dari user
	//handler : mapping input dari user -> struct input
	//services : melakukan mapping dari struct input ke struct
	//repository
	//db
}
