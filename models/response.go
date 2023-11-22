package models

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int64       `json:"Code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data,omitempty"`
}

func HandleResponse(ctx *gin.Context, r *Response) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.JSON(int(r.Code), r)
}
