package controllers

import (
	"bank-app-server/src/db"
	"bank-app-server/src/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	// Kullanıcıdan hesap numarasını alın (URL parametresi olarak)
	accountNumber := mux.Vars(r)["accountNumber"]

	var account models.Account
	if err := db.DB.Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		http.Error(w, "Hesap bulunamadı", http.StatusNotFound)
		return
	}

	// Hesap bakiyesini döndür
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{
		"balance": account.Balance,
	})
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	// Transfer bilgilerini alalım
	var transfer struct {
		FromAccount   string  `json:"from_account"`
		ToAccount     string  `json:"to_account"`
		Amount        float64 `json:"amount"`
		Description   string  `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	var fromAccount, toAccount models.Account

	// Gönderen ve alıcı hesaplarını sorgula
	if err := db.DB.Where("account_number = ?", transfer.FromAccount).First(&fromAccount).Error; err != nil {
		http.Error(w, "Gönderen hesap bulunamadı", http.StatusNotFound)
		return
	}
	if err := db.DB.Where("account_number = ?", transfer.ToAccount).First(&toAccount).Error; err != nil {
		http.Error(w, "Alıcı hesap bulunamadı", http.StatusNotFound)
		return
	}

	// Gönderenin yeterli bakiyesi var mı?
	if fromAccount.Balance < transfer.Amount {
		http.Error(w, "Yetersiz bakiye", http.StatusBadRequest)
		return
	}

	// Transfer işlemini gerçekleştir
	fromAccount.Balance -= transfer.Amount
	toAccount.Balance += transfer.Amount

	// Hesap bakiyelerini güncelle
	db.DB.Save(&fromAccount)
	db.DB.Save(&toAccount)

	// İşlem kaydı oluştur
	transaction := models.Transaction{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        transfer.Amount,
		Description:   transfer.Description,
	}
	db.DB.Create(&transaction)

	// Başarı mesajı
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transfer başarılı",
	})
}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	// Rastgele hesap numarası üret
	account.AccountNumber = fmt.Sprintf("%08d", rand.Intn(100000000))
	account.Balance = 0.0 // Başlangıçta sıfır bakiye

	// Kullanıcıyı Context'ten al
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Error(w, "Kullanıcı bilgisi alınamadı", http.StatusInternalServerError)
		return
	}

	// Hesap ile kullanıcıyı ilişkilendir
	account.UserID = user.ID

	// Hesabı veritabanına kaydet
	if err := db.DB.Create(&account).Error; err != nil {
		http.Error(w, "Hesap oluşturulamadı", http.StatusInternalServerError)
		return
	}

	// Başarı mesajı
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func ListAccountsHandler(w http.ResponseWriter, r *http.Request) {
	// Kullanıcıyı Context'ten al
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Error(w, "Kullanıcı bilgisi alınamadı", http.StatusInternalServerError)
		return
	}

	// Kullanıcının tüm hesaplarını sorgula
	var accounts []models.Account
	if err := db.DB.Where("user_id = ?", user.ID).Find(&accounts).Error; err != nil {
		http.Error(w, "Hesaplar bulunamadı", http.StatusInternalServerError)
		return
	}

	// Hesapları döndür
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func AccountSummaryHandler(w http.ResponseWriter, r *http.Request) {
	// Kullanıcıyı Context'ten al
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Error(w, "Kullanıcı bilgisi alınamadı", http.StatusInternalServerError)
		return
	}

	// Kullanıcının hesaplarını sorgula
	var accounts []models.Account
	if err := db.DB.Where("user_id = ?", user.ID).Find(&accounts).Error; err != nil {
		http.Error(w, "Hesaplar bulunamadı", http.StatusInternalServerError)
		return
	}

	// Son birkaç işlemi sorgula
	var transactions []models.Transaction
	if err := db.DB.Limit(10).Where("from_account_id IN ? OR to_account_id IN ?", accountIDsFromAccounts(accounts), accountIDsFromAccounts(accounts)).Order("created_at desc").Find(&transactions).Error; err != nil {
		http.Error(w, "İşlem geçmişi alınamadı", http.StatusInternalServerError)
		return
	}

	// Hesapların bakiyesi ve son işlemleri JSON olarak döndür
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"accounts":     accounts,
		"transactions": transactions,
	})
}

// accountIDsFromAccounts bir dizi Account alır ve onların ID'lerini döndürür
func accountIDsFromAccounts(accounts []models.Account) []uint {
	var accountIDs []uint
	for _, account := range accounts {
		accountIDs = append(accountIDs, account.ID)
	}
	return accountIDs
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	var deposit struct {
		AccountNumber string  `json:"account_number"`
		Amount        float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	var account models.Account
	if err := db.DB.Where("account_number = ?", deposit.AccountNumber).First(&account).Error; err != nil {
		http.Error(w, "Hesap bulunamadı", http.StatusNotFound)
		return
	}

	// Bakiyeyi artır
	account.Balance += deposit.Amount
	if err := db.DB.Save(&account).Error; err != nil {
		http.Error(w, "Bakiyeye para yatırılırken hata oluştu", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Para başarıyla yatırıldı",
	})
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	var withdraw struct {
		AccountNumber string  `json:"account_number"`
		Amount        float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&withdraw); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	var account models.Account
	if err := db.DB.Where("account_number = ?", withdraw.AccountNumber).First(&account).Error; err != nil {
		http.Error(w, "Hesap bulunamadı", http.StatusNotFound)
		return
	}

	// Yetersiz bakiye kontrolü
	if account.Balance < withdraw.Amount {
		http.Error(w, "Yetersiz bakiye", http.StatusBadRequest)
		return
	}

	// Bakiyeyi azalt
	account.Balance -= withdraw.Amount
	if err := db.DB.Save(&account).Error; err != nil {
		http.Error(w, "Bakiyeden para çekilirken hata oluştu", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Para başarıyla çekildi",
	})
}