# Aplikasi Todo-list dengan Role-Based Access Control (RBAC) menggunakan Go dan MySQL

Aplikasi ini adalah sebuah API Todo-list yang dibangun dengan menggunakan bahasa pemrograman **Go** dan database **MySQL**. Aplikasi ini menerapkan **Role-Based Access Control (RBAC)** untuk mengelola hak akses pengguna. Hanya pengguna dengan peran **Editor** yang dapat melakukan operasi CRUD (Create, Read, Update, Delete) pada Todo-list, sementara pengguna dengan peran **Admin** dapat mengelola data pengguna tanpa dapat mengakses Todo-list.

## 📋 Fitur

- **Autentikasi JWT**: Mengamankan endpoint API dengan JSON Web Tokens.
- **Role-Based Access Control**: Mengatur hak akses berdasarkan peran pengguna (Admin dan Editor).
- **CRUD Users**: Hanya Admin yang dapat membuat, membaca, memperbarui, dan menghapus pengguna.
- **CRUD Todos**: Hanya Editor yang dapat membuat, membaca, memperbarui, dan menghapus todo.
- **Keamanan Password**: Password pengguna disimpan dengan hashing menggunakan bcrypt.
- **Logging dan Recovery**: Middleware untuk logging request dan recovery dari panic.

## 🛠 Teknologi yang Digunakan

- **Bahasa Pemrograman**: Go (Golang)
- **Framework Web**: Echo
- **Database**: MySQL
- **Autentikasi**: JSON Web Tokens (JWT)
- **ORM/Driver**: `github.com/go-sql-driver/mysql`
- **Password Hashing**: `golang.org/x/crypto/bcrypt`

## 🔧 Prasyarat

Sebelum memulai, pastikan Anda telah menginstal hal-hal berikut di sistem Anda:

- [Go](https://golang.org/dl/) (versi 1.16 atau lebih baru)
- [MySQL](https://dev.mysql.com/downloads/mysql/)
- [Git](https://git-scm.com/downloads)

## 📦 Instalasi

### 1. Clone Repository

Clone repository ini ke direktori lokal Anda:

```bash
git clone https://github.com/adtbch/CRUD-Role-Golang.git
cd repository-name
```

### 2. Instal Dependensi

Jalankan perintah berikut untuk mengunduh semua dependensi yang diperlukan:

```bash
go mod tidy
```

### 3. Konfigurasi Database

#### a. Buat Database dan Pengguna di MySQL

1. **Login ke MySQL**:

    ```bash
    mysql -u root -p
    ```

2. **Buat Database**:

    ```sql
    CREATE DATABASE tododb;
    ```

3. **Buat Pengguna dan Berikan Hak Akses**:

    ```sql
    CREATE USER 'todouser'@'localhost' IDENTIFIED BY 'yourpassword';
    GRANT ALL PRIVILEGES ON tododb.* TO 'todouser'@'localhost';
    FLUSH PRIVILEGES;
    ```

    *Ganti `yourpassword` dengan password yang Anda inginkan.*

#### b. Sesuaikan String Koneksi di `main.go`

Buka file `main.go` dan sesuaikan variabel `dsn` dengan kredensial MySQL Anda:

```go
dsn := "todouser:yourpassword@tcp(127.0.0.1:3306)/tododb?parseTime=true"
```

*Ganti `yourpassword` dengan password yang Anda tetapkan sebelumnya.*

## 🚀 Menjalankan Aplikasi

Setelah semua konfigurasi selesai, jalankan aplikasi dengan perintah berikut:

```bash
go run main.go
```

Jika berhasil, Anda akan melihat output seperti berikut:

```
⇨ http server started on [::]:8080
```

## 📄 Dokumentasi API

### 1. Autentikasi

#### **Login**

- **Endpoint**: `POST /login`
- **Deskripsi**: Mengautentikasi pengguna dan menghasilkan token JWT.
- **Body**:

    ```json
    {
        "username": "admin",
        "password": "admin123"
    }
    ```

- **Respons**:

    ```json
    {
        "token": "your_jwt_token"
    }
    ```

### 2. CRUD Users (Admin Only)

#### **Mendapatkan Semua Pengguna**

- **Endpoint**: `GET /users`
- **Deskripsi**: Mengambil daftar semua pengguna.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    ```

- **Respons**:

    ```json
    [
        {
            "id": 1,
            "username": "admin",
            "role": "Admin"
        },
        {
            "id": 2,
            "username": "editor1",
            "role": "Editor"
        }
    ]
    ```

#### **Membuat Pengguna Baru**

- **Endpoint**: `POST /users`
- **Deskripsi**: Membuat pengguna baru.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    Content-Type: application/json
    ```

- **Body**:

    ```json
    {
        "username": "editor2",
        "password": "password123",
        "role": "Editor"
    }
    ```

- **Respons**:

    ```json
    {
        "message": "User created successfully"
    }
    ```

#### **Memperbarui Pengguna**

- **Endpoint**: `PUT /users/{id}`
- **Deskripsi**: Memperbarui informasi pengguna.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    Content-Type: application/json
    ```

- **Body**:

    ```json
    {
        "username": "editor2",
        "password": "newpassword123",
        "role": "Editor"
    }
    ```

- **Respons**:

    ```json
    {
        "message": "User updated successfully"
    }
    ```

#### **Menghapus Pengguna**

- **Endpoint**: `DELETE /users/{id}`
- **Deskripsi**: Menghapus pengguna berdasarkan ID.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    ```

- **Respons**:

    ```json
    {
        "message": "User deleted successfully"
    }
    ```

### 3. CRUD Todos (Editor Only)

#### **Mendapatkan Semua Todo untuk Pengguna Terkait**

- **Endpoint**: `GET /todos`
- **Deskripsi**: Mengambil daftar semua todo untuk pengguna yang sedang login.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    ```

- **Respons**:

    ```json
    [
        {
            "id": 1,
            "title": "Membeli bahan makanan",
            "completed": false,
            "user_id": 2
        },
        {
            "id": 2,
            "title": "Mengerjakan tugas",
            "completed": true,
            "user_id": 2
        }
    ]
    ```

#### **Membuat Todo Baru**

- **Endpoint**: `POST /todos`
- **Deskripsi**: Membuat todo baru.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    Content-Type: application/json
    ```

- **Body**:

    ```json
    {
        "title": "Belajar Go"
    }
    ```

- **Respons**:

    ```json
    {
        "message": "Todo created successfully"
    }
    ```

#### **Memperbarui Todo**

- **Endpoint**: `PUT /todos/{id}`
- **Deskripsi**: Memperbarui informasi todo.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    Content-Type: application/json
    ```

- **Body**:

    ```json
    {
        "title": "Belajar Go Lang",
        "completed": true
    }
    ```

- **Respons**:

    ```json
    {
        "message": "Todo updated successfully"
    }
    ```

#### **Menghapus Todo**

- **Endpoint**: `DELETE /todos/{id}`
- **Deskripsi**: Menghapus todo berdasarkan ID.
- **Headers**:

    ```
    Authorization: Bearer your_jwt_token
    ```

- **Respons**:

    ```json
    {
        "message": "Todo deleted successfully"
    }
    ```

## 🧪 Pengujian

Anda dapat menggunakan alat seperti **Postman** atau **cURL** untuk menguji endpoint API.

### **Contoh Penggunaan dengan Postman**

1. **Login untuk Mendapatkan Token JWT**
    - **Method**: `POST`
    - **URL**: `http://localhost:8080/login`
    - **Body**: Raw JSON

        ```json
        {
            "username": "admin",
            "password": "admin123"
        }
        ```

    - **Respons**:

        ```json
        {
            "token": "your_jwt_token"
        }
        ```

2. **Mengakses Endpoint dengan Token**
    - **Headers**:
        ```
        Authorization: Bearer your_jwt_token
        Content-Type: application/json
        ```

    - **Endpoint Lainnya**: Sesuaikan dengan dokumentasi API di atas.

### **Menggunakan cURL**

#### **Login**

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "username": "admin",
    "password": "admin123"
}'
```

#### **Membuat Todo Baru (Editor Only)**

```bash
curl -X POST http://localhost:8080/todos \
-H "Authorization: Bearer your_jwt_token" \
-H "Content-Type: application/json" \
-d '{
    "title": "Belajar Go"
}'
```

## 📚 Struktur Proyek

```
.
├── main.go
├── models.go
├── dto.go
├── handler.go
├── routes.go
├── role_middleware.go
├── go.mod
└── go.sum
```

- **main.go**: File utama yang menginisialisasi aplikasi dan koneksi database.
- **models.go**: Definisi struktur data (User, Todo).
- **dto.go**: Data Transfer Objects untuk menerima dan mengirim data.
- **handler.go**: Fungsi handler untuk setiap endpoint API.
- **routes.go**: Konfigurasi rute dan penerapan middleware.
- **role_middleware.go**: Middleware untuk mengelola hak akses berdasarkan peran.
- **go.mod** & **go.sum**: File modul Go untuk mengelola dependensi.


## 📫 Kontak

Jika Anda memiliki pertanyaan atau masukan, silakan hubungi:

- **Email**: adit.bachtiar091@gmail.com

---

Terima kasih telah menggunakan aplikasi ini! Semoga bermanfaat untuk kebutuhan Anda.
```