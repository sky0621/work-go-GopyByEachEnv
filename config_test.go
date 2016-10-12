package GopyByEachEnv

import "testing"

func TestReadConfig(t *testing.T) {
	readConfig2(ReadConfigParam{targetFile: "cmd/config.json"})
	if config2 == nil {
		t.Fatal("config2がnil")
	}

	// [MEMO] t.Log() では何もコンソールに出力されなかったので t.Error() で試しに json 読み込んだ結果を出力。
	// for _, spec := range config2.CopySpecs {
	// 	t.Error(spec.CopyFrom.FromDir)
	// 	t.Error(spec.CopyFrom.FromFile)
	// 	for _, copyTo := range spec.CopyTos {
	// 		t.Error(copyTo.ToDir)
	// 		t.Error(copyTo.ToFile)
	// 		for _, rep := range copyTo.Replaces {
	// 			t.Error(rep.ReplaceFrom)
	// 			t.Error(rep.ReplaceTo)
	// 		}
	// 	}
	// }
}
