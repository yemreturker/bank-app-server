
# Proje Dosya Yapısı

Bu dosya, projenin dosya yapısını ve her dosyanın ne işe yaradığını açıklamaktadır.

## Bank-App-Server Dosya Yapısı:

```
bank-app-server/
│
├── src/
│   ├── cmd/
│   │   └── main.go                      # Ana giriş noktası
│   ├── controllers/
│   │   ├── account_controller.go        # Hesap ile ilgili işlemler
│   │   ├── auth_controller.go           # Kimlik doğrulama ve oturum işlemleri
│   │   ├── transaction_controller.go    # İşlem geçmişi ve transferler
│   │   └── scheduled_transfer_controller.go  # Planlı transferler
│   ├── db/
│   │   └── database.go                  # Veritabanı bağlantısı ve migration
│   ├── middlewares/
│   │   └── auth_middleware.go           # JWT tabanlı kimlik doğrulama
│   ├── models/
│   │   ├── account.go                   # Hesap modeli
│   │   ├── transaction.go               # İşlem modeli
│   │   ├── user.go                      # Kullanıcı modeli
│   │   └── scheduled_transfer.go        # Planlı transfer modeli
│   └── utils/
│       └── helpers.go                   # Yardımcı fonksiyonlar (örneğin, accountIDsFromAccounts)
│   ├── go.mod                           # Go modülü ve bağımlılıklar
│   ├── go.sum                           # Go modüllerinin checksum bilgileri
│
├── bank.db                              # SQLite veritabanı dosyası (veritabanı migration sonrası otomatik oluşur)
├── .gitignore                           # Versiyon kontrolünde gereksiz dosyaları dışarıda tutmak için
├── README.md                            # Projenin genel açıklaması ve kullanım talimatları
└── PROJECT_STRUCTURE.md                 # Proje dosya yapısını açıklayan bu dosya
```