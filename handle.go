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

// /gopy/start の無限ループ操作フラグ
var continueServer = true

func watchStart(w http.ResponseWriter, r *http.Request) {
	log.Println("監視開始")
	io.WriteString(w, "監視中。。。\n")

	var baseTime = time.Now()

	log.Println(config.CopySpecs)
	// [MEMO] configはグローバルにせざるを得ないか。。。 WebServer起動からハンドラーにパラメータ渡しできればいいんだけど。
	for _, spec := range config.CopySpecs {
		copyFrom := spec.CopyFrom.FromDir + spec.CopyFrom.FromFile
		log.Println("[コピー元] " + copyFrom)
		modifier := Modifier{beforeTime: baseTime, checkFilePath: copyFrom}
		go copyExec(w, modifier, copyFrom, spec.CopyTos)
	}
}

func copyExec(w http.ResponseWriter, modifier Modifier, copyFrom string, copyTos []CopyTo) {
	for continueServer {
		doCopy, err := modifier.isModify()
		if err != nil {
			log.Println(err)
			io.WriteString(w, "監視対象ファイルの更新日チェック時にエラー発生\n")
			continueServer = false
			return
		}
		if doCopy {
			copier := Copier{copyFrom: copyFrom, copyTos: copyTos}
			// [MEMO] fluentパターン採用。Goでも、この使い方、ありなのか・・・？
			copier.copyToEachDir().replaceEachToFile().renameTmp()
			if copier.err != nil {
				log.Println(copier.err)
				io.WriteString(w, "監視対象ファイルのコピー時にエラー発生\n")
				continueServer = false
				return
			}
		}
		time.Sleep(time.Duration(args.SleepSecond) * time.Second)
	}
}

func watchEnd(w http.ResponseWriter, r *http.Request) {
	log.Println("監視終了")
	io.WriteString(w, "監視終了\n")
	continueServer = false
}
