package GopyByEachEnv

import (
	"reflect"
	"testing"
)

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

	expected := buildExpected()
	if !reflect.DeepEqual(config, expected) {
		t.Fatalf("結果と期待値が要素レベルで異なっています。\n[結果]：%#v\n[期待値]：%#v", config, expected)
	}
}

func buildExpected() *Config {
	cfg := Config{
		[]CopySpec{
			CopySpec{
				CopyFrom{
					FromDir:  "_sample/from/",
					FromFile: "sample.properties",
				},
				[]CopyTo{
					CopyTo{
						ToDir:  "_sample/to01/",
						ToFile: "sample.properties",
						Replaces: []Replace{
							Replace{
								ReplaceFrom: "returnPath=returnPath01@example.jp",
								ReplaceTo:   "returnPath=convertedPath01@example.jp",
							},
							Replace{
								ReplaceFrom: "textTemplate.encoding=UTF-8",
								ReplaceTo:   "textTemplate.encoding=SJIS",
							},
						},
					},
					CopyTo{
						ToDir:  "_sample/to02/",
						ToFile: "sampleB.properties",
					},
				},
			},
			CopySpec{
				CopyFrom{
					FromDir:  "_sample/from/",
					FromFile: "sample2.properties",
				},
				[]CopyTo{
					CopyTo{
						ToDir:  "_sample/to01/",
						ToFile: "sample2b.properties",
						Replaces: []Replace{
							Replace{
								ReplaceFrom: "subject=メール２",
								ReplaceTo:   "subject=メール２(コンバート)",
							},
						},
					},
				},
			},
		},
	}
	return &cfg
}
