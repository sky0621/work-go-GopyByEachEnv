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
	if err != nil {
		log.Println(err)
		c.err = err
		return c
	}
	defer from.Close()

	for _, copyTo := range c.copyTos {
		to, toErr := os.Create(copyTo.ToDir + copyTo.ToFile)
		if toErr != nil {
			log.Println(toErr)
			c.err = toErr
			return c
		}
		defer to.Close()

		_, cpErr := io.Copy(to, from)
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
		if scnErr != nil {
			log.Println(scnErr)
			c.err = scnErr
			return c
		}
	}
	return c
}

func (c *Copier) renameTmp() {
	// FIXME replaceEachToFile()のデバッグが終わったら実装！
}

// // [MEMO] とりあえず作った版にしても、雑さが・・・
// func renameTmp() {
// 	var err error
// 	toFiles := config.m["to"]
// 	toArray := strings.Split(toFiles, "$")
//
// 	for _, to := range toArray {
// 		toSplit := strings.Split(to, "\t")
// 		err = os.Rename(toSplit[0]+".tmp", toSplit[0])
// 		if err != nil {
// 			log.Println(err)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 	}
// }
