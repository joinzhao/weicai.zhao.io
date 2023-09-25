package tools

import "github.com/golang-jwt/jwt/v4"

const signedKey = "asdefsasdasdasdw"

func JwtSign(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signedKey))
}
func JwtParse(tokenString string, claims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signedKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}
