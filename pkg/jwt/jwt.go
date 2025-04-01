package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Secret string
}

type JWTData struct {
	Phone string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) GenerateToken(phone string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) ParseToken(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	phone := t.Claims.(jwt.MapClaims)["phone"]
	return t.Valid, &JWTData{
		Phone: phone.(string),
	}
}
