package GopyByEachEnv

import (
	"log"
	"net/http"
)

// アプリの終了コード
const (
	ExitCodeOK          int = iota // 0
	ExitCodeConfigError            // 1
	ExitCodeCopyError              // 2
	ExitCodePanic                  // 3
)

// ExitCode ... デフォルトOK
var ExitCode = ExitCodeOK

var config *Config

// Exec ... エラー時は終了コードを上書き！
func Exec() {

	// [MEMO] こうすればパニック起きても、ここでハンドリングできる？
	defer func() {
		if err := recover(); err != nil {
			log.Printf("パニック発生！リカバー！　Error: %s", err)
			ExitCode = ExitCodePanic
		}
	}()

	config = ParseConfig("config.json")
	if config == nil {
		ExitCode = ExitCodeConfigError
		return
	}

	// [MEMO] 同じhttpでもServeMuxというのもある様子。。。
	// [MEMO] ルーティングはサードパーティ使うのがよいみたいだけど、このツールではコピー監視の開始・終了だけだからこれで十分
	http.HandleFunc("/gopy/", gopyHandler)

	log.Println("コピー元ファイルの監視を開始します：", args.Port)
	// [MEMO] graceful というパッケージが書籍に紹介されてるけど、なんかハンドラーに至らず落ちるので諦め。
	if err := http.ListenAndServe(args.Port, nil); err != nil {
		log.Println(err)
		ExitCode = ExitCodeCopyError
	}
}
