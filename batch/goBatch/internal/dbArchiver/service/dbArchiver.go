package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/kazGear/portfolio/goBatch/pkg/constants"
)

type DbArchiver struct {
	BackupDir  string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewDbArchiverService() *DbArchiver {
	return &DbArchiver{
		BackupDir:  "/app/backup",
		DBHost:     os.Getenv("DB_HOST"), // db
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func (b *DbArchiver) Archive() error {
	date := time.Now().Format(constants.DateTime)
	backupFile := filepath.Join(
		b.BackupDir,
		fmt.Sprintf("kaz_app_%s.dump", date),
	)

	if err := os.MkdirAll(b.BackupDir, 0755); err != nil {
		return err
	}

	fmt.Println("PostgreSQL backup started.")

	file, err := os.Create(backupFile)
	if err != nil {
		return fmt.Errorf("create backup file: %w", err)
	}
	defer file.Close()

// tmp start .........................................................
path, err := exec.LookPath("pg_dump")
if err != nil {
	return fmt.Errorf("pg_dump not found: %w", err)
}

version, err := exec.Command("pg_dump", "--version").CombinedOutput()
if err != nil {
	return fmt.Errorf("pg_dump version check failed: %w", err)
}

fmt.Printf("pg_dump path: %s\n", path)
fmt.Printf("pg_dump version: %s", version)
// tmp end .........................................................

	// コマンドを作成
	cmd := exec.Command(
		"pg_dump",
		"-h", b.DBHost,
		"-U", b.DBUser,
		"-d", b.DBName,
		"-Fc",
	)

	cmd.Env = append(os.Environ(), "PGPASSWORD="+b.DBPassword)

	var stderr bytes.Buffer
	cmd.Stdout = file
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		_ = os.Remove(backupFile)

		return fmt.Errorf(
			"pg_dump failed: %w: %s",
			err,
			stderr.String(),
		)
	}
	fmt.Printf("Backup created: %s\n", backupFile)

	if err := deleteOldBackup(b.BackupDir); err != nil {
		return err
	}
	fmt.Println("Old backups deleted.")

	return nil
}

func deleteOldBackup(backupDir string) error {
	// バックアップファイル群
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return err
	}

	// ファイル削除のラインを設定
	threshold := time.Now().AddDate(0, 0, -30)

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "kaz_app_") {
			continue
		}

		// ファイル情報取得
		info, err := entry.Info()
		if err != nil {
			return err
		}

		// ラインより古い変更日か
		if info.ModTime().Before(threshold) {
			// 古いファイルを削除
			if err := os.Remove(filepath.Join(backupDir, entry.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}