package controllers

import (
	"bank-app-server/src/db"
	"bank-app-server/src/models"
	"encoding/json"
	"net/http"
)

func TransactionHistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Kullanıcıyı Context'ten al
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Error(w, "Kullanıcı bilgisi alınamadı", http.StatusInternalServerError)
		return
	}

	// Kullanıcının tüm hesaplarını bul
	var accounts []models.Account
	if err := db.DB.Where("user_id = ?", user.ID).Find(&accounts).Error; err != nil {
		http.Error(w, "Hesaplar bulunamadı", http.StatusInternalServerError)
		return
	}

	// Hesapların ID'lerini çek
	var accountIDs []uint
	for _, account := range accounts {
		accountIDs = append(accountIDs, account.ID)
	}

	// Hesaplara bağlı tüm işlemleri sorgula
	var transactions []models.Transaction
	if err := db.DB.Where("from_account_id IN ? OR to_account_id IN ?", accountIDs, accountIDs).Find(&transactions).Error; err != nil {
		http.Error(w, "İşlem geçmişi bulunamadı", http.StatusInternalServerError)
		return
	}

	// İşlem geçmişini döndür
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}