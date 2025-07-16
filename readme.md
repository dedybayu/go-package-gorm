# 📦 Go Package GORM

Latihan dasar penggunaan **Golang** dengan struktur project terpisah (modular) dan integrasi **GORM** sebagai ORM (Object Relational Mapping) untuk koneksi ke database MySQL.

## ✨ Fitur

- Struktur project yang rapi menggunakan package `model`, `config`, dan `middleware`.
- Koneksi database menggunakan GORM.
- Migrasi otomatis tabel `users`.
- Contoh routing HTTP menggunakan Gorilla Mux.
- Middleware logging request.

---

## 📂 Struktur Project

```
go-package-gorm/
├── config/             # Konfigurasi database dan environment
│   └── db.go
├── middleware/         # Middleware HTTP (logging)
│   └── logger.go
├── model/              # Struct model (GORM)
│   └── user.go
├── go.mod
├── go.sum
└── main.go             # Entry point aplikasi
```

---

## 🚀 Instalasi & Menjalankan

### 1. Clone Repository

```bash
git clone https://github.com/dedybayu/go-package-gorm.git
cd go-package-gorm
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Atur Konfigurasi Database

Edit file `config/db.go` dan sesuaikan dengan kredensial MySQL kamu:

```go
dsn := "root:password@tcp(127.0.0.1:3306)/namadb?charset=utf8mb4&parseTime=True&loc=Local"
```

### 4. Jalankan Aplikasi

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

---

## 🥪 Endpoint Contoh

- `GET /hello`\
  Mengembalikan JSON:
  ```json
  { "ok": true }
  ```

---

## 🧰 Tools yang Digunakan

- [Golang](https://golang.org/)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Logrus](https://github.com/sirupsen/logrus)

---

## 📝 Catatan

- Migrasi tabel `users` dijalankan otomatis pada startup melalui `AutoMigrate`.
- Pastikan database target **sudah dibuat** terlebih dahulu sebelum menjalankan aplikasi.

---

## 📌 Lisensi

Project ini dibuat untuk keperluan latihan dan tidak memiliki lisensi resmi.

---

## 🙇‍♂️ Kontributor

- [Dedy Bayu Setiawan](https://github.com/dedybayu)

