package repositories

import (
	"github/ersa97/med-record/models"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Login(ctx *gin.Context, req models.LoginRequest) (*models.Response, error)
}
