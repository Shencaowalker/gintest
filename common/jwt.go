package common

import (
	"gintest1/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

//创建声明结构体
type Claims struct {
	UserId uint
	jwt.StandardClaims
}


//创建token
func ReleaseToken(user model.User)(string,error){
	expirationTime := time.Now().Add(7*24*time.Hour) //设置过期时间
	//创建声明 
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "shencao",
			Subject: "user token",
		},
	}
	//基于声明生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	//根据自定义jwtKey密钥签署token
	tokenString,err := token.SignedString(jwtKey)

	if err !=nil{
		return "",err
	}
	return tokenString,err
}


//解析token并返回参数
func ParseToken(tokenString string)(*jwt.Token,*Claims,error){
	claims :=&Claims{}
	token,err := jwt.ParseWithClaims(tokenString,claims,func(token *jwt.Token)(i interface{},err error){
		return jwtKey,nil
	})
	return token,claims,err
	
}
