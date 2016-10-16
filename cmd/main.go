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

// 【心得群】
// ・panicは最終手段。無暗に使わない。通常は多値の一部として error を返し、愚直でも呼び元でエラーチェック！
// ・正規表現のパフォーマンスが悪いので、極力 strings パッケージを使う！
// ・mapは避け、構造体を使う！（※mapに対する操作はスレッドセーフでない。排他制御にはsync.RWMutexを使う。）
// ・reflectパッケージによるリフレクションは避ける！（黒魔術）
// ・巨大なstructの生成・継承はしない！
// ・並行処理を使いすぎない！（多くの場所では直列で十分。ホットスポットでのみゴルーチンを使う。）
// ・Goのコード自体を読む！

// ・積極的にpath/filepathを使う！
