package jwt_

import (
	"time"

	"github.com/almazatun/lien-court/pkg/common/helper"
	jtoken "github.com/golang-jwt/jwt/v5"
)

func GenToken(email, id string) (t string, err error) {
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	tn, err := token.SignedString([]byte(helper.GetEnvVar("JWT_SECRET_KEY")))

	if err != nil {
		return "", nil
	}

	return tn, nil
}
