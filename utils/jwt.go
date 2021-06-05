package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey string
var expminute int

func init() {
	secretKey = "qwcvo0321vfqxas"
	expminute = 3600 * 24
}

type UserClaims struct {
	PWD  string `json:"pwd"`
	Name string `json:"name"`
	// jwt.StandardClaims        //嵌套了这个结构体就实现了Claim接口
}

func JwtInit(key string, exp int) {
	secretKey = key
	expminute = exp
}

func JwtCreate(username, pwd string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["pwd"] = pwd
	claims["username"] = username
	//过期时间
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expminute)).Unix()
	claims["nbf"] = time.Now().Unix()
	claims["iat"] = time.Now().Unix()

	token.Claims = claims

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}
}

func JwtParse(tokenss string) (user *UserClaims, err error) {
	user = &UserClaims{}
	token, err := jwt.Parse(tokenss, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	user.PWD = claim["pwd"].(string)
	user.Name = claim["username"].(string)
	return
}
