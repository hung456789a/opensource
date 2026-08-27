package tokenuntil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func CreateAcessToken(user User, secret string, expiry int) string {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := user.JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user User, secret string, expiry int) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claimsRefresh := JwtCustomRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func IsAuthorized() {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string error) {
	token, err := jwt.Parse(requesToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Error("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["id"].(string), nil
}

func main() {
	mySigningKey := []byte("AllYourBase")
	fmt.Println(mySigningKey)
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}
	claims := JwtCustomClaims{
		Name: "John Doe",
		ID:   "123",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24))},
	}
	fmt.Println(claims)
}
