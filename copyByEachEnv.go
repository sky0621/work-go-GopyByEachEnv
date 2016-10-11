package GopyByEachEnv

import (
	"log"
	"net/http"
	"os"
)

// Exec ... プログラムのメインロジック
func Exec() {
	// [MEMO] 同じhttpでもServeMuxというのもある様子。。。
	// [MEMO] ルーティングはサードパーティ使うのがよいみたいだけど、このツールではコピー監視の開始・終了だけだからこれで十分
	http.HandleFunc("/gopy/", gopyHandler)

	log.Println("コピー元ファイルの監視を開始します：", args.Port)
	// [MEMO] graceful というパッケージが書籍に紹介されてるけど、なんかハンドラーに至らず落ちるので諦め。
	if err := http.ListenAndServe(args.Port, nil); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("コピー元ファイルの監視を終了します：", args.Port)
}
