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

// [MEMO] グローバル・・・。嫌だったら、httpハンドラー内に渡すには、（リクエストごとにゴルーチンらしいから）チャネル使う？
var config *Config

// Exec ... エラー時は終了コードを上書き！
func Exec() {

	// [MEMO] こうすればパニック起きても、ここでハンドリングできる？
	// [MEMO] httpハンドラー内がリクエストごとにゴルーチン実行らしく、その中でパニック起きても、ここに飛んでこない・・・。パニックそもそも起こすべきじゃないかもだけど、ちょっと検討必要。
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
	// [MEMO] ツールの性質上、httpハンドラー使う必要まったくなしだったけど、単に勉強のためだけに使った。が、複数リクエスト受けられ、そのつどゴルーチン実行（？）は、むしろデバッグしんどいだけか・・・。
	// [MEMO] 今回はこの仕組みにしたけど、単純なツールは単純につくるべき。WebAPI関連のをつくるときにhttpハンドラー使えばいいね。
	http.HandleFunc("/gopy/", gopyHandler)

	log.Println("コピー元ファイルの監視を開始します：", args.Port)
	// [MEMO] graceful というパッケージが書籍に紹介されてるけど、なんかハンドラーに至らず落ちるので諦め。
	if err := http.ListenAndServe(args.Port, nil); err != nil {
		log.Println(err)
		ExitCode = ExitCodeCopyError
	}
}
