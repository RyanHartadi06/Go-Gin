package handler

import (
	"bwastartup/mahasiswa"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mahasiswaHandler struct {
	mahasiswaService mahasiswa.Service
}

func NewMahasiswaHandler(mahasiswaService mahasiswa.Service) *mahasiswaHandler {
	return &mahasiswaHandler{mahasiswaService: mahasiswaService}
}

func (h *mahasiswaHandler) Register(c *gin.Context){
	var input mahasiswa.RegisterMahasiswaInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest , nil)
	}
	mahasiswa , err := h.mahasiswaService.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest , nil)
	}
	c.JSON(http.StatusOK , mahasiswa)
}