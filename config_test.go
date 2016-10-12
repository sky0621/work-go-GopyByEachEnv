package GopyByEachEnv

import "testing"

const configSpec string = "[仕様] cmd直下のconfig.jsonの解析ができる"

func TestReadConfig(t *testing.T) {
	t.Log(configSpec)
	config := ParseConfig("cmd/config.json")
	if config == nil {
		t.Fatal("結果がnil")
	}
	if len(config.CopySpecs) < 1 {
		t.Fatal("結果の要素数が０以下")
	}

}
