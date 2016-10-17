package main

import (
	"flag"
	"log"
	"os"

	GopyByEachEnv "github.com/sky0621/work-go-GopyByEachEnv"
)

var version = "0.1.0"

// [MEMO] 書籍では「cmd」ディレクトリの下にプロジェクト名を表すディレクトリを作り、その下に main.go を置く事例が多いとあったが・・・。
func main() {
	args := parseFlag()
	if args == nil {
		return
	}
	GopyByEachEnv.Exec(args)
	log.Println("ExitCode: ", GopyByEachEnv.ExitCode)
	os.Exit(GopyByEachEnv.ExitCode)
}

// [MEMO] flag使ってる場合のテストコードって、どうやって書くんだ？
func parseFlag() *GopyByEachEnv.Args {
	var showVersion bool
	var port string
	var sleep int
	// [MEMO] プログラム引数の扱いも便利なサードパーティがある様子。go-flags とか。
	flag.BoolVar(&showVersion, "v", false, "バージョン")
	flag.StringVar(&port, "p", "7109", "接続先ポート")
	flag.IntVar(&sleep, "s", 10, "コピー元ファイルの変更監視間隔(秒)")
	flag.Parse()
	if showVersion {
		log.Println("version:", version)
		return nil
	}
	return &GopyByEachEnv.Args{Port: ":" + port, SleepSecond: sleep}
}

// 【心得群】
// ・panicは最終手段。無暗に使わない。通常は多値の一部として error を返し、愚直でも呼び元でエラーチェック！
// ・正規表現のパフォーマンスが悪いので、極力 strings パッケージを使う！
// ・mapは避け、構造体を使う！（※mapに対する操作はスレッドセーフでない。排他制御にはsync.RWMutexを使う。）
// ・reflectパッケージによるリフレクションは避ける！（黒魔術）
// ・巨大なstructの生成・継承はしない！
// ・並行処理を使いすぎない！（多くの場所では直列で十分。ホットスポットでのみゴルーチンを使う。）
// ・Goのコード自体を読む！

// ・積極的にpath/filepathを使う！（※物理的なファイルを扱う場合は「path」ではなく「path/filepath」パッケージを使わないと、ディレクトリトラバーサル等のセキュリティリスクとなる）
// ・積極的にdeferを使う！（※リソース解放漏れを防ぐ）
// ・積極的にUTF-8を使う！
