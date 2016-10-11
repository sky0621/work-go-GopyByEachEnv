package main

import (
	"log"
	"os"

	GopyByEachEnv "github.com/sky0621/work-go-GopyByEachEnv"
)

// [MEMO] 書籍では「cmd」ディレクトリの下にプロジェクト名を表すディレクトリを作り、その下に main.go を置く事例が多いとあったが・・・。
func main() {
	GopyByEachEnv.Exec()
	log.Println("ExitCode: " + string(GopyByEachEnv.ExitCode))
	os.Exit(GopyByEachEnv.ExitCode)
}
