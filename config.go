package GopyByEachEnv

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// [MEMO] argsとconfigまではアプリ起動時に読み込むでOKかな。
// [MEMO] config変数もconfig.goファイル内にあれば副作用関数でもテストコードは書きやすいかな。
var config *Config
var config2 *Config2

// Config ... コピー元ファイルやコピー先などの設定
type Config struct {
	m map[string]string // [MEMO] 構造体の意味なし。。。当初作った config.txt の構造が雑なままとりあえず構造体に持ち込んだだけ
}

// Config2 ... コピー元ファイルやコピー先などの設定
type Config2 struct {
	CopySpecs []CopySpec
}

// CopySpec ...
type CopySpec struct {
	CopyFrom CopyFrom
	CopyTos  []CopyTo
}

// CopyFrom ...
type CopyFrom struct {
	FromDir  string
	FromFile string
}

// CopyTo ...
type CopyTo struct {
	ToDir    string
	ToFile   string
	Replaces []Replace
}

// Replace ...
type Replace struct {
	ReplaceFrom string
	ReplaceTo   string
}

// ReadConfigParam ...
type ReadConfigParam struct {
	targetFile string
}

func readConfig2(p ReadConfigParam) {
	file, err := ioutil.ReadFile(p.targetFile)
	if err != nil {
		log.Println("config.jsonがファイル自体読み込めない！")
		log.Println(err)
		ExitCode = ExitCodeConfigError
		return
	}
	jsonErr := json.Unmarshal(file, &config2)
	if jsonErr != nil {
		log.Println(jsonErr)
		ExitCode = ExitCodeConfigError
		return
	}
}

// [MEMO] 雑すぎ・・・。
func readConfig() {
	var fp *os.File
	// [MEMO] main.goの階層からの相対パス
	fp, err := os.OpenFile("config.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Println(err)
		ExitCode = ExitCodeConfigError
		return
	}

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
	err = scanner.Err()
	if err != nil {
		log.Println(err)
		ExitCode = ExitCodeConfigError
		return
	}

	config = &Config{m: m}
}
