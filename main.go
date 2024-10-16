// main.go
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key") // Ganti dengan kunci rahasia Anda

func main() {
	// Koneksi ke MySQL
	dsn := "root@tcp(127.0.0.1:3306)/tododb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	// Membuat tabel users
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(255) UNIQUE NOT NULL,
            password_hash VARCHAR(255) NOT NULL,
            role VARCHAR(50) NOT NULL
        )`)
	if err != nil {
		panic("Failed to create users table: " + err.Error())
	}

	// Membuat tabel todos
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS todos (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            completed BOOLEAN NOT NULL,
            user_id INT,
            FOREIGN KEY (user_id) REFERENCES users(id)
        )`)
	if err != nil {
		panic("Failed to create todos table: " + err.Error())
	}

	// Insert Admin user if not exists
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password: " + err.Error())
	}

	// Check if admin user exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "admin").Scan(&count)
	if err != nil {
		panic("Failed to check for admin user: " + err.Error())
	}

	if count == 0 {
		_, err = db.Exec("INSERT INTO users (username, password_hash, role) VALUES (?, ?, ?)",
			"admin", string(hashedPassword), "Admin")
		if err != nil {
			panic("Failed to insert admin user: " + err.Error())
		}
	}

	// Initialize Echo
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Provide DB connection to handlers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// Set up routes
	Route(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
