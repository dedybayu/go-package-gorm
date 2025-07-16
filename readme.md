<div align="right">
  <a href="#en">🇺🇸 English</a> | <a href="#id">🇮🇩 Bahasa Indonesia</a>
</div>


# 📦 Go Package GORM (Bahasa Indonesia) <a name="id"></a>

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

Copy file `.env_example` menjadi `.env`:  
Edit file `.env` dan sesuaikan dengan kredensial MySQL kamu.

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
- `GET /users`\
  Mengembalikan JSON:
  ```json
  [
    {
        "id": 1,
        "name": "Name1",
        "email": "email1",
        "password": "pass1",
        "created_at": "2025-07-16T22:08:52+07:00",
        "updated_at": "2025-07-16T22:08:52+07:00"
    },
    {
        "id": 2,
        "name": "Name2",
        "email": "email2",
        "password": "pass2",
        "created_at": "2025-07-16T22:11:14+07:00",
        "updated_at": "2025-07-16T22:11:14+07:00"
    }
  ]
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


---

# 📦 Go Package GORM (English Version) <a name="en"></a>

A basic practice project using **Golang** with a modular structure and **GORM** as the Object Relational Mapping (ORM) to connect to a MySQL database.

## ✨ Features

- Well-organized project structure with `model`, `config`, and `middleware` packages
- MySQL database connection using GORM
- Automatic migration for the `users` table
- HTTP routing using Gorilla Mux
- Logging middleware for incoming requests

---

## 📂 Project Structure

```
go-package-gorm/
├── config/             # Database and environment configuration
│   └── db.go
├── middleware/         # HTTP middleware (logging)
│   └── logger.go
├── model/              # GORM data model
│   └── user.go
├── go.mod
├── go.sum
└── main.go             # Application entry point
```

---

## 🚀 Installation & Run

### 1. Clone Repository

```bash
git clone https://github.com/dedybayu/go-package-gorm.git
cd go-package-gorm
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Database

Copy `.env_example` to `.env`:  
Edit the `.env` file and update your MySQL credentials accordingly.

### 4. Run the Application

```bash
go run main.go
```

The server will run at `http://localhost:8080`

---

## 🥪 Sample Endpoints

- `GET /hello`\
  Returns JSON:
  ```json
  { "ok": true }
  ```
- `GET /users`\
  Returns JSON:
  ```json
  [
    {
        "id": 1,
        "name": "Name1",
        "email": "email1",
        "password": "pass1",
        "created_at": "2025-07-16T22:08:52+07:00",
        "updated_at": "2025-07-16T22:08:52+07:00"
    },
    {
        "id": 2,
        "name": "Name2",
        "email": "email2",
        "password": "pass2",
        "created_at": "2025-07-16T22:11:14+07:00",
        "updated_at": "2025-07-16T22:11:14+07:00"
    }
  ]
  ```

---

## 🧰 Tools Used

- [Golang](https://golang.org/)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Logrus](https://github.com/sirupsen/logrus)

---

## 📝 Notes

- The `users` table is automatically migrated on application startup using `AutoMigrate`.
- Ensure the target database **already exists** before starting the application.

---

## 📌 License

This project is intended for learning purposes and is not officially licensed.

---

## 🙇‍♂️ Contributor

- [Dedy Bayu Setiawan](https://github.com/dedybayu)

