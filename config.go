package GopyByEachEnv

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// [MEMO] サブパッケージモジュールで使う機会が発生するかもなので、以降の構造体はexported

// Config ... コピー元ファイルやコピー先などの設定
type Config struct {
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

// ParseConfig ...
// [MEMO] main.go起点とテストコード起点とでカレント変わるため config.json 読むためのパスも変わる。それぞれで引数渡さず、main.go起点はデフォルトにしたい。
// [MEMO] ↑でも、デフォルト引数とかオーバーロードは Go には無いみたい・・・。
func ParseConfig(targetFilePath string) *Config {
	file, err := ioutil.ReadFile(targetFilePath)
	if err != nil {
		log.Printf("%s の読み込みに失敗しました。指定のパスにファイルが存在するか確認してください。 [ERROR] %s", targetFilePath, err)
		return nil
	}

	var config *Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Printf("%s のJSONとしての解析に失敗しました。指定のファイルがJSONとして正しい形式か確認してください。 [ERROR] %s", targetFilePath, err)
		return nil
	}
	// log.Println(config)
	return config
}
