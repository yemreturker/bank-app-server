package controllers

import (
	"bank-app-server/src/db"
	"bank-app-server/src/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	// Şifreyi hashle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Şifre hash'leme hatası", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Kullanıcıyı veritabanına ekle
	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Kullanıcı oluşturulamadı", http.StatusInternalServerError)
		return
	}

	// Başarı mesajı
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Kullanıcı başarıyla kaydedildi",
	})
}

var jwtKey = []byte("your_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		http.Error(w, "Kullanıcı bulunamadı", http.StatusUnauthorized)
		return
	}

	// Şifre kontrolü
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Yanlış şifre", http.StatusUnauthorized)
		return
	}

	// JWT Token oluştur
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Token oluşturulamadı", http.StatusInternalServerError)
		return
	}

	// Token'ı döndür
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func PasswordResetHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	// Gelen veriyi çözümlüyoruz
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	// Kullanıcıyı Context'ten alıyoruz
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Error(w, "Kullanıcı bulunamadı", http.StatusInternalServerError)
		return
	}

	// Mevcut şifreyi kontrol ediyoruz
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		http.Error(w, "Eski şifre yanlış", http.StatusUnauthorized)
		return
	}

	// Yeni şifreyi hash'leyip güncelliyoruz
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Şifre güncellenemedi", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	if err := db.DB.Save(&user).Error; err != nil {
		http.Error(w, "Şifre güncellenemedi", http.StatusInternalServerError)
		return
	}

	// Başarı mesajı döndür
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Şifre başarıyla güncellendi",
	})
}

func TokenRefreshHandler(w http.ResponseWriter, r *http.Request) {
	// Kullanıcıyı Context'ten alıyoruz
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Error(w, "Kullanıcı bulunamadı", http.StatusInternalServerError)
		return
	}

	// Yeni token oluşturuyoruz
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Token oluşturulamadı", http.StatusInternalServerError)
		return
	}

	// Yeni token'ı döndür
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}