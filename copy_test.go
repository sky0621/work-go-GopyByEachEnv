package GopyByEachEnv

import (
	"os"
	"testing"
)

const copySpec string = "[仕様] 任意のパスにあるファイルが任意のパスにコピーされる"

func TestCopyToEachDir(t *testing.T) {
	t.Log(copySpec)
	copyTos := []CopyTo{CopyTo{ToDir: "testdata/", ToFile: "copyTo.txt"}}
	copier := Copier{copyFrom: "testdata/copyFrom.txt", copyTos: copyTos}
	copier.copyToEachDir()
	if copier.err != nil {
		t.Fatalf("コピー時にエラー発生 %s", copier.err)
	}
	// [MEMO] 後始末用のメソッドを定義する機能はないのかな。
	os.Remove("testdata/copyTo.txt")
}
