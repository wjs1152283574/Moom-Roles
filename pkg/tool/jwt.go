package tool

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/it-moom/moom-roles/pkg/errors"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// NewJWT 新建一个jwt实例
func NewJWT(screct string) *JWT {
	return &JWT{
		[]byte(screct),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(userId, issuer string, tll int64) (string, error) {

	// 过期时间处理
	expiresAt := time.Now().Add(time.Second * time.Duration(tll)).Unix()
	// 构造Payload
	claims := jwt.StandardClaims{
		ExpiresAt: expiresAt,         // 过期时间
		Id:        SimpleUUID(),      // Jti:token唯一标识符
		IssuedAt:  time.Now().Unix(), // 颁发时间
		Issuer:    issuer,            // 颁发机构
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   userId,            // 主题
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	// token有误
	if err != nil {
		return nil, errors.ErrInvalidToken()
	}
	// token校验成功
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.ErrInvalidToken()
}

// RefreshToken 刷新token
func (j *JWT) RefreshToken(tokenString, screct string) (string, error) {

	// 解析token
	claims, err := NewJWT(screct).ParseToken(tokenString)

	// token无效
	if err != nil {
		return "", errors.ErrInvalidToken()
	}

	// 重新生成token
	return j.CreateToken(claims.Subject, claims.Issuer, claims.ExpiresAt)
}
