package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service 
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	//tangkap input
	//map input dari user ke struct ke RegisterInput
	//struct diatas kita passing sbg params service 
	var input user.RegisterUserInput 
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}

		response := helper.APIResponse("Register failed" , http.StatusUnprocessableEntity , "error" , errorMessage)
		c.JSON(http.StatusUnprocessableEntity , response)
		return 
	}

	newUser , err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register failed" , http.StatusBadRequest , "error" , err.Error())
		c.JSON(http.StatusBadRequest , response)
		return
	}

	formatter := user.FormatUser(newUser , "tokentokentoken")
	response := helper.APIResponse("Account has been registered" , http.StatusOK , "success" , formatter)
	c.JSON(http.StatusOK , response)
	
}

func (h *userHandler) LoginUser(c *gin.Context){
		//user memasukkan input email dan password
		//input ditangkapn handler
		//mapping dari input user ke input struct
		//input struct passing ke service
		//di service find dg bantuin repository user dengan email x
		//mencocokkan password
		var input user.LoginInput
		err := c.ShouldBindJSON(&input)
		if err != nil {
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors" : errors}
	
			response := helper.APIResponse("Login failed" , http.StatusUnprocessableEntity , "error" , errorMessage)
			c.JSON(http.StatusUnprocessableEntity , response)
			return 
		}

		loggedInUser , err := h.userService.Login(input)

		if err != nil {
			errorMessage := gin.H{"errors" : err.Error()}
			response := helper.APIResponse("Login failed" , http.StatusUnprocessableEntity , "error" , errorMessage)
			c.JSON(http.StatusUnprocessableEntity , response)
			return 
		}

		formatter := user.FormatUser(loggedInUser , "tokentokentoken")
		response := helper.APIResponse("Login Successfuly" , http.StatusOK , "success" , formatter)
		c.JSON(http.StatusOK , response)
}