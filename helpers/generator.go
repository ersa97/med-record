package helpers

import (
	"github/ersa97/med-record/data"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	key []byte
}

func NewJWT(key []byte) JWT {
	return JWT{
		key: key,
	}
}

func (j JWT) GenerateJWTToken(user data.User, permission []string) (jwtToken string, err error) {

	timeExpired, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	expired := time.Now().Add(time.Hour * time.Duration(timeExpired)).Format("2006-01-02 15:04:05")

	claims := make(jwt.MapClaims)
	claims["Id"] = user.ID
	claims["Username"] = user.Username
	claims["RoleId"] = user.Roleid
	claims["Permission"] = permission
	claims["Expired"] = expired

	jwtToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.key)
	return
}
