package AuthService

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	jwt.StandardClaims
	UserID int
}

func CreateToken(userId int) (string, error) {
	claim := jwtClaim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func CreateCookie(userId int) http.Cookie {
	jwt, err := CreateToken(userId)
	if err != nil {
		panic(err)
	}

	cookie := http.Cookie{
		Name:     "Auth",
		Value:    jwt,
		Path:     "/",
		Domain:   "http://localhost:3000",
		Expires:  time.Now().Local().Add(1 * time.Hour),
		MaxAge:   int(time.Hour),
		Secure:   true,
		HttpOnly: true,
	}

	return cookie
}

func CreateEmptyCookie() http.Cookie {
	cookie := http.Cookie{
		Name:     "Auth",
		Value:    "",
		Path:     "/",
		Domain:   "http://localhost:3000",
		Expires:  time.Now(),
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	}

	return cookie
}
