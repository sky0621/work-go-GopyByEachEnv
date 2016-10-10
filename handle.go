package GopyByEachEnv

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func gopyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w, r)
	default:
		log.Println("GETのみ受け付けます")
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

func watchStart(w http.ResponseWriter, r *http.Request) {
	log.Println("監視開始")
	io.WriteString(w, "監視中。。。")

	var baseTime time.Time
	for {
		var nowTime = modTime()
		if baseTime != nowTime {
			// [MEMO] 以下３関数は共通のインタフェースでも被せてコマンドパターン化すべきか？
			// [MEMO] 中でやってること似てる部分あると思うけど、継承のないGoではテンプレートメソッドパターン使いたい時どうする？
			copyToEachDir()
			replaceEachToFile()
			renameTmp()
			baseTime = nowTime
		}

		time.Sleep(time.Duration(args.SleepSecond) * time.Second)
	}
}

func watchEnd(w http.ResponseWriter, r *http.Request) {
	log.Println("監視終了")
	// [MEMO] ここでOSイグジットは乱暴？
	os.Exit(0)
}

func modTime() time.Time {
	fromDir := config.m["fromDir"]
	fromFile := config.m["fromFile"]
	from := fromDir + fromFile
	fp, err := os.Stat(from)
	handleErr(err)
	return fp.ModTime()
}
