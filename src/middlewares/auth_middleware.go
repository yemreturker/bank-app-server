package middlewares

import (
	"bank-app-server/src/db"
	"bank-app-server/src/models"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Yetkisiz erişim", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Yetkisiz erişim", http.StatusUnauthorized)
			return
		}

		// Kullanıcıyı veritabanından bulalım ve Context'e ekleyelim
		var user models.User
		if err := db.DB.Where("username = ?", claims.Subject).First(&user).Error; err != nil {
			http.Error(w, "Kullanıcı bulunamadı", http.StatusUnauthorized)
			return
		}

		// Context'e kullanıcıyı ekle
		ctx := context.WithValue(r.Context(), "user", &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}