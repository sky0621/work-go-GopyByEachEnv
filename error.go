package GopyByEachEnv

import (
	"log"
	"os"
)

// [MEMO] エラーハンドリングのセオリー調べよう。。。　Javaで言うとアプリ固有の例外規定するけど、Goでは？　パニックの使いどころは？
func handleErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
