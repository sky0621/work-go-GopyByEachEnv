package GopyByEachEnv

import (
	"log"
	"net/http"
)

// [MEMO] 自明と思われるものでも何かしらコメント書かないとgolintでWARN出るのは、けっこうきつい。
// アプリの終了コード
const (
	ExitCodeOK          int = iota // 0
	ExitCodeError                  // 1
	ExitCodeConfigError            // 2
	ExitCodeCopyError              // 3
	ExitCodePanic                  // 4
)

// ExitCode ... 終了コード上書き用にデフォルトをOKでセット
var ExitCode = ExitCodeOK

// Exec ... エラー時は終了コードを上書き！
func Exec() {
	// [MEMO] こうすればパニック起きても、ここでハンドリングできる？
	defer func() {
		log.Println("パニックチェック開始")
		if err := recover(); err != nil {
			log.Println(err)
			ExitCode = ExitCodePanic
		}
	}()

	// [MEMO] 同じhttpでもServeMuxというのもある様子。。。
	// [MEMO] ルーティングはサードパーティ使うのがよいみたいだけど、このツールではコピー監視の開始・終了だけだからこれで十分
	http.HandleFunc("/gopy/", gopyHandler)

	log.Println("コピー元ファイルの監視を開始します：", args.Port)
	// [MEMO] graceful というパッケージが書籍に紹介されてるけど、なんかハンドラーに至らず落ちるので諦め。
	if err := http.ListenAndServe(args.Port, nil); err != nil {
		log.Println(err)
		ExitCode = ExitCodeError
	}
	log.Println("コピー元ファイルの監視を終了します：", args.Port)
}
