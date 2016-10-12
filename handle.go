package GopyByEachEnv

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func gopyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w, r)
	default:
		io.WriteString(w, "GETのみ受け付けます")
	}
}

// [MEMO] Goではメッセージ管理はどうするんだろう。
const usage = "/gopy/start ないし /gopy/end のみ受け付けます。（/gopy/start/ のように最後に「/」つけてもダメ）"

func handleGet(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	// [MEMO] ドキュメントルートの「/」の前も１要素としてSplitされ、空文字としてカウントされる(?)ようなので、要素数３でチェック
	if len(segs) != 3 {
		io.WriteString(w, usage)
		return
	}

	switch segs[2] {
	case "start":
		watchStart(w, r)
	case "end":
		watchEnd(w, r)
	default:
		io.WriteString(w, usage)
	}
}

// [MEMO] /gopy/end が叩かれたら false セットして /gopy/start の無限ループを抜けさせる方法にしよう
var continueServer = true

func watchStart(w http.ResponseWriter, r *http.Request) {
	log.Println("監視開始")
	io.WriteString(w, "監視中。。。")

	var baseTime = time.Now()
	// [MEMO] configはグローバルにせざるを得ないか。。。 WebServer起動からハンドラーにパラメータ渡しできればいいんだけど。
	for _, spec := range config.CopySpecs {
		modifier := Modifier{
			beforeTime:    baseTime,
			checkFilePath: spec.CopyFrom.FromDir + spec.CopyFrom.FromFile}
		for continueServer {
			doCopy, err := modifier.isModify()
			if err != nil {
				log.Println(err)
				return
			}
			if doCopy {
				log.Println("<<<<< DO COPY >>>>>")
				// [MEMO] 以下３関数は共通のインタフェースでも被せてコマンドパターン化すべきか？
				// [MEMO] 中でやってること似てる部分あると思うけど、継承のないGoではテンプレートメソッドパターン使いたい時どうする？
				// copyToEachDir()
				// replaceEachToFile()
				// renameTmp()
			}
			time.Sleep(time.Duration(args.SleepSecond) * time.Second)
		}
	}
}

func watchEnd(w http.ResponseWriter, r *http.Request) {
	log.Println("監視終了")
	continueServer = false
}
