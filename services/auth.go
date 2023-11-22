package services

import (
	"context"
	"errors"
	"github/ersa97/med-record/data"
	"github/ersa97/med-record/helpers"
	"github/ersa97/med-record/models"
	"github/ersa97/med-record/repositories"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *pgx.Conn
}

func NewAuthService() repositories.AuthRepository {
	return &AuthService{}
}

func (a *AuthService) Login(ctx *gin.Context, req models.LoginRequest) (res *models.Response, err error) {

	queries := data.New(a.DB)
	user, err := queries.GetUserByUsername(context.Background(), req.Username)
	if err != nil {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: "wrong username or password",
		}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return &models.Response{
			Code:    http.StatusUnauthorized,
			Message: "wrong username or password",
		}, errors.New("wrong username or password")
	}

	_, err = queries.GetRoleById(context.Background(), user.Roleid)
	if err != nil {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: "role not found",
		}, errors.New("role not found")
	}

	permissionskey, err := queries.GetPermissionkeysByRoleId(context.Background(), user.Roleid)
	if err != nil {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: "permission not found",
		}, errors.New("permission not found")
	}

	var permission []string

	for _, v := range permissionskey {
		if v.Valid {
			permission = append(permission, v.String)
		}
	}

	jwtKey := helpers.NewJWT([]byte(os.Getenv("JWT_KEY")))

	token, err := jwtKey.GenerateJWTToken(user, permission)

	return &models.Response{
		Code:    http.StatusOK,
		Message: "login successful",
		Data: models.ResponseLogin{
			UserId: user.ID,
			Token:  token,
		},
	}, nil
}
