
# API Dökümantasyonu

Bu döküman, Bank-App-Server projesindeki tüm API endpoint'lerini detaylı bir şekilde açıklar.

## Kimlik Doğrulama ve Kullanıcı Yönetimi

### 1. Kullanıcı Kaydı (Signup)
- **Endpoint**: `/signup`
- **Metod**: `POST`
- **Açıklama**: Yeni kullanıcı kaydı yapar.
- **Girdi**:
  ```json
  {
    "username": "kullaniciadi",
    "password": "sifre"
  }
  ```
- **Çıktı**:
  ```json
  {
    "message": "Kullanıcı başarıyla kaydedildi"
  }
  ```

### 2. Kullanıcı Girişi (Login)
- **Endpoint**: `/login`
- **Metod**: `POST`
- **Açıklama**: Kullanıcı giriş yapar ve JWT token alır.
- **Girdi**:
  ```json
  {
    "username": "kullaniciadi",
    "password": "sifre"
  }
  ```
- **Çıktı**:
  ```json
  {
    "token": "jwt_token"
  }
  ```

### 3. Şifre Sıfırlama (Password Reset)
- **Endpoint**: `/password/reset`
- **Metod**: `POST`
- **Açıklama**: Kullanıcı şifresini sıfırlar.
- **Girdi**:
  ```json
  {
    "old_password": "eski_sifre",
    "new_password": "yeni_sifre"
  }
  ```
- **Çıktı**:
  ```json
  {
    "message": "Şifre başarıyla güncellendi"
  }
  ```

## Hesap Yönetimi

### 4. Hesap Oluşturma
- **Endpoint**: `/account`
- **Metod**: `POST`
- **Açıklama**: Yeni bir hesap oluşturur.
- **Girdi**:
  ```json
  {}
  ```
- **Çıktı**:
  ```json
  {
    "account_number": "12345678",
    "balance": 0.0
  }
  ```

### 5. Hesap Bakiyesi Sorgulama
- **Endpoint**: `/balance/{accountNumber}`
- **Metod**: `GET`
- **Açıklama**: Belirtilen hesap numarasına ait bakiyeyi döndürür.
- **Çıktı**:
  ```json
  {
    "balance": 500.00
  }
  ```

### 6. Hesap Listeleme
- **Endpoint**: `/accounts`
- **Metod**: `GET`
- **Açıklama**: Kullanıcıya ait tüm hesapları listeler.
- **Çıktı**:
  ```json
  [
    {
      "account_number": "12345678",
      "balance": 500.00
    },
    {
      "account_number": "87654321",
      "balance": 1000.00
    }
  ]
  ```

### 7. Hesap Özeti (Account Summary)
- **Endpoint**: `/account/summary`
- **Metod**: `GET`
- **Açıklama**: Kullanıcının hesaplarını ve son işlemlerini özetler.
- **Çıktı**:
  ```json
  {
    "accounts": [...],
    "transactions": [...]
  }
  ```

### 8. Para Yatırma (Deposit)
- **Endpoint**: `/account/deposit`
- **Metod**: `POST`
- **Açıklama**: Belirtilen hesaba para yatırır.
- **Girdi**:
  ```json
  {
    "account_number": "12345678",
    "amount": 100.00
  }
  ```
- **Çıktı**:
  ```json
  {
    "message": "Para başarıyla yatırıldı"
  }
  ```

### 9. Para Çekme (Withdraw)
- **Endpoint**: `/account/withdraw`
- **Metod**: `POST`
- **Açıklama**: Belirtilen hesaptan para çeker.
- **Girdi**:
  ```json
  {
    "account_number": "12345678",
    "amount": 50.00
  }
  ```
- **Çıktı**:
  ```json
  {
    "message": "Para başarıyla çekildi"
  }
  ```

## Para Transferi

### 10. Para Transferi (Transfer)
- **Endpoint**: `/transfer`
- **Metod**: `POST`
- **Açıklama**: Bir hesaptan başka bir hesaba para transferi yapar.
- **Girdi**:
  ```json
  {
    "from_account": "12345678",
    "to_account": "87654321",
    "amount": 100.00,
    "description": "Transfer açıklaması"
  }
  ```
- **Çıktı**:
  ```json
  {
    "message": "Transfer başarılı"
  }
  ```

## İşlem Geçmişi

### 11. İşlem Geçmişi (Transaction History)
- **Endpoint**: `/transactions`
- **Metod**: `GET`
- **Açıklama**: Kullanıcının tüm işlem geçmişini listeler.
- **Çıktı**:
  ```json
  [
    {
      "ID": 1,
      "FromAccountID": 2,
      "ToAccountID": 3,
      "Amount": 100.00,
      "Description": "Hesaplar arası transfer",
      "CreatedAt": "2023-09-17T12:34:56Z"
    },
    ...
  ]
  ```

## Planlı Transferler

### 12. Planlı Transfer Ekleme (Scheduled Transfer)
- **Endpoint**: `/transfer/scheduled`
- **Metod**: `POST`
- **Açıklama**: Gelecekte gerçekleştirilecek bir planlı transfer oluşturur.
- **Girdi**:
  ```json
  {
    "from_account": "12345678",
    "to_account": "87654321",
    "amount": 150.00,
    "scheduled_date": "2023-12-25T10:00:00Z"
  }
  ```
- **Çıktı**:
  ```json
  {
    "message": "Planlı transfer başarıyla eklendi"
  }
  ```
