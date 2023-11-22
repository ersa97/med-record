package controllers

import (
	"github/ersa97/med-record/models"
	"github/ersa97/med-record/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Controller struct {
	DB *sqlx.DB
}

func (c *Controller) Login(ctx *gin.Context) {

	var req models.LoginRequest

	if err := ctx.BindJSON(&req); err != nil {
		models.HandleResponse(ctx, &models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		models.HandleResponse(ctx, &models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	res, err := services.NewAuthService().Login(ctx, req)
	if err != nil {
		models.HandleResponse(ctx, &models.Response{
			Code:    res.Code,
			Message: res.Message,
		})
		return
	}
	models.HandleResponse(ctx, &models.Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
	})
}
