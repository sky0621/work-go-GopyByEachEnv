package GopyByEachEnv

import "testing"

const (
	SPEC string = "[仕様：cmd直下のconfig.jsonの解析ができる] "
)

func TestReadConfig(t *testing.T) {
	config := ParseConfig("cmd/config.json")
	if config == nil {
		t.Fatal(SPEC + "[事象：結果がnil]")
	}
	if len(config.CopySpecs) < 1 {
		t.Fatal(SPEC + "[事象：結果の要素数が０以下]")
	}

}
