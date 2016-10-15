package GopyByEachEnv

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Copier ...
type Copier struct {
	copyFrom string
	copyTos  []CopyTo // [MEMO] 中途半端に依存。config関連はI knowというよりEveryone knowという感じだから良しか。
	err      error    // [MEMO] fluentにして、かつ、errorを返すにはこうなるけど、ありかな・・・。
}

// [MEMO] インタフェース作って構造体をコピー用、置換用、リネーム用切り替えるか、今のfluentパターンか。
// [MEMO] 厳密にやると、それぞれの用途で使うconfig要素が違うから構造体切り替えがいいのかも。
// [MEMO] 今回は、コピー先要素を使うという部分は共通しているから、fluentパターンにしてみよう。でも、エラーを返す都合上、副作用発生。

// [MEMO] やっぱりロジック似通ってる。TemplateMethodパターン使いたい。

func (c *Copier) copyToEachDir() *Copier {
	if c.err != nil {
		return c
	}
	from, err := os.Open(c.copyFrom)
	log.Println("[コピー元] " + from.Name())
	if err != nil {
		log.Println(err)
		c.err = err
		return c
	}
	defer from.Close()

	for _, copyTo := range c.copyTos {
		from.Seek(0, 0) // 毎回、読み込み位置をリセットしておかないと同一ファイルからの読み込みは２回目以降サイズ０

		to, toErr := os.Create(copyTo.ToDir + copyTo.ToFile)
		log.Println("[コピー先] " + to.Name())
		if toErr != nil {
			log.Println(toErr)
			c.err = toErr
			return c
		}
		defer to.Close()

		size, cpErr := io.Copy(to, from)
		log.Printf("コピー実行サイズ： %d\n", size)
		if cpErr != nil {
			log.Println(cpErr)
			c.err = cpErr
			return c
		}
	}
	return c
}

func (c *Copier) replaceEachToFile() *Copier {
	if c.err != nil {
		return c
	}
	for _, copyTo := range c.copyTos {
		replaceFrom, err := os.Open(copyTo.ToDir + copyTo.ToFile)
		if err != nil {
			log.Println(err)
			c.err = err
			return c
		}
		defer replaceFrom.Close()

		replaceTo, toErr := os.Create(copyTo.ToDir + copyTo.ToFile + ".tmp") // [MEMO] TempFile関数使うべきかな。
		if toErr != nil {
			log.Println(toErr)
			c.err = toErr
			return c
		}
		defer replaceTo.Close()

		scanner := bufio.NewScanner(replaceFrom)
		writer := bufio.NewWriter(replaceTo)
		for scanner.Scan() {
			line := scanner.Text()
			for _, replace := range copyTo.Replaces {
				isMatch, _ := regexp.MatchString(replace.ReplaceFrom, line)
				if isMatch {
					line = strings.Replace(line, replace.ReplaceFrom, replace.ReplaceTo, -1)
				}
			}
			writer.WriteString(line + "\r\n")
		}
		writer.Flush()

		scnErr := scanner.Err()
		if handleErr(c, scnErr) {
			return c
		}
	}
	return c
}

func (c *Copier) renameTmp() *Copier {
	if c.err != nil {
		return c
	}
	for _, copyTo := range c.copyTos {
		err := os.Rename(copyTo.ToDir+copyTo.ToFile+".tmp", copyTo.ToDir+copyTo.ToFile)
		if handleErr(c, err) {
			return c
		}
	}
	return c
}

func handleErr(c *Copier, e error) bool {
	if e == nil {
		return false
	}
	log.Println(e)
	c.err = e
	return true
}
