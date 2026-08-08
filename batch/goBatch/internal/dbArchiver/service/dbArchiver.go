package service

import (
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

	// dumpの出力先となる空ファイルを作成
	file, err := os.Create(backupFile)

	if err != nil {
		return fmt.Errorf("Create backup file: %w", err)
	}
	defer file.Close()

	// コマンドを作成
	cmd := exec.Command(
		"pg_dump",
		"-h", b.DBHost,
		"-U", b.DBUser,
		"-d", b.DBName,
		"-Fc",
	)

	// 既存環境変数に環境変数を追加
	cmd.Env = append(os.Environ(), "PGPASSWORD=" + b.DBPassword)
	// pg_dumpの標準出力を先ほど作ったファイルに接続する
	cmd.Stdout = file

	// 作成したコマンドを実行
	if err := cmd.Run(); err != nil {
		_ = os.Remove(backupFile)
		return fmt.Errorf("pg_dump failed: %w", err)
	}

	fmt.Printf("Backup created: %s\n", backupFile)

	// バックアップファイル群
	entries, err := os.ReadDir(b.BackupDir)
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
			if err := os.Remove(filepath.Join(b.BackupDir, entry.Name())); err != nil {
				return err
			}
		}
	}

	fmt.Println("Old backups deleted.")

	return nil
}
