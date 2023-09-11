package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mohamedsamara/go-vue/constants"
	"github.com/mohamedsamara/go-vue/models"
	"github.com/mohamedsamara/go-vue/utils"
	"gorm.io/gorm"
)

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func CreateAccessToken(ttl time.Duration, id *uuid.UUID) (string, error) {
	claims := &jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * ttl).Unix(),
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func WithJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt")

		token, err := VerifyJWT(tokenString)

		if err != nil {
			utils.WritePermissionDenied(w)
			return
		}

		if !token.Valid {
			utils.WritePermissionDenied(w)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		userId := claims["id"]

		if !ok {
			utils.WritePermissionDenied(w)
			return
		}

		db := r.Context().Value(constants.DB_CONTEXT).(*gorm.DB)

		var user models.User

		db.First(&user, "id = ?", userId)

		if user.ID.String() != userId {
			utils.WritePermissionDenied(w)
			return
		}

		userContext := utils.FormatUserResponse(user)

		ctx := context.WithValue(r.Context(), constants.AUTH_CONTEXT, userContext)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
