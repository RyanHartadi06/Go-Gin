package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		var errors []string

		for _, e := range err.(validator.ValidationErrors){
			errors = append(errors, e.Error())
		}

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