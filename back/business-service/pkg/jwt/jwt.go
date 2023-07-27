package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"valtec/pkg/config"
)

var key = []byte(`JVC+L6TgyzMXLJOekPtXfyfSZDWkOdvt+hU2SASs1y0=`)

type Payload struct {
	claims jwt.MapClaims
}

func New() *Payload {
	p := &Payload{
		claims: map[string]interface{}{},
	}
	p.claims["exp"] = time.Now().Add(time.Duration(config.GetIntConfigSevere("jwt", "exp")) * time.Second).Unix()
	return p
}

func (p *Payload) Set(key string, value any) *Payload {
	p.claims[key] = value
	return p
}

func (p *Payload) Token() string {
	tok, err := jwt.NewWithClaims(jwt.SigningMethodHS256, p.claims).SignedString(key)
	if err != nil {
		panic("获取token失败: " + err.Error())
	}
	return tok
}

func Parse(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token, please login again")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return &Payload{claims: claims}, nil
}

func (p *Payload) Get(key string) any {
	return p.claims[key]
}
