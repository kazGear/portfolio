package service

import (
	"fmt"
	"os/exec"
)

type DbArchiver struct {}

func NewDbArchiverService() *DbArchiver {
	return &DbArchiver{}
}

func (b *DbArchiver) Archive() error {
	filepath := "/home/kazuki/app/portfolio/batch/forLinux/BackupPostgres.sh"
	cmd := exec.Command("sh", filepath)

	// 標準出力と標準エラー出力をまとめて取得
	output, err := cmd.CombinedOutput()

	if err != nil {
		// shスクリプトが exit 1 などで終了した場合
		return fmt.Errorf("DB archiver script error: %w\n%s", err, output)
	}
	return nil
}
