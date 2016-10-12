package GopyByEachEnv

import "flag"

const defaultPort = "7109"
const defaultSleepSec = 10 // 10秒置きにコピー元ファイルの変更を監視

// [MEMO] Goではvalは無いんだっけ？ Goでもやっぱりグローバルは極力避けるものなのか？ とはいえ、Scalaのように純粋関数を考慮するものでもない？
var args = parseFlag()

// Args ... プログラム引数
type Args struct {
	Port        string
	SleepSecond int
}

func parseFlag() *Args {
	var port string
	var sleep int
	// [MEMO] プログラム引数の扱いも便利なサードパーティがある様子。go-flags とか。
	flag.StringVar(&port, "p", defaultPort, "接続先ポート")
	flag.IntVar(&sleep, "s", defaultSleepSec, "コピー元ファイルの変更監視間隔(秒)")
	flag.Parse()
	return &Args{Port: port, SleepSecond: sleep}
}
