package GopyByEachEnv

import (
	"bufio"
	"os"
	"strings"
)

// [MEMO] argsとconfigまではアプリ起動時に読み込むでOKかな。
var config = readConfig()

// Config ... コピー元ファイルやコピー先などの設定
type Config struct {
	m map[string]string // [MEMO] 構造体の意味なし。。。当初作った config.txt の構造が雑なままとりあえず構造体に持ち込んだだけ
}

// [MEMO] 雑すぎ・・・。
func readConfig() *Config {
	var fp *os.File
	// [MEMO] main.goの階層からの相対パス
	fp, err := os.OpenFile("config.txt", os.O_RDONLY, 0)
	handleErr(err)

	defer fp.Close()

	m := map[string]string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		// [MEMO] 設定ファイルの読み込み系のお決まりロジックないしサードパーティ探さないと。。。　そもそもconfig.txtの構造が適当
		kv := strings.Split(line, "=")
		v, ok := m[kv[0]]
		if ok {
			t := v
			m[kv[0]] = t + "$" + kv[1]
		} else {
			m[kv[0]] = kv[1]
		}
	}
	handleErr(scanner.Err())

	return &Config{m: m}
}
