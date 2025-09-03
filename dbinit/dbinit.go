package dbinit

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
	"time"
)

// Init database files

// Connect will establish database connection and return a handle
func Connect(hostname string, port uint16, username, password, database string, ssl bool) (*sqlx.DB, error) {
	sslMode := "disable"
	if ssl {
		sslMode = "require"
	}
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		hostname,
		port,
		username,
		password,
		database,
		sslMode)

	var err error
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(32)
	db.SetMaxIdleConns(32)
	db.SetConnMaxLifetime(time.Duration(1800) * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func Populate(db *sqlx.DB, filepath string) error {
	if filepath == "" {
		return fmt.Errorf("filepath is empty")
	}

	b, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}
	sqlText := strings.TrimSpace(string(b))
	if sqlText == "" {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	if _, err := db.ExecContext(ctx, sqlText); err != nil {
		return fmt.Errorf("failed to execute SQL: %w", err)
	}
	return nil
}
