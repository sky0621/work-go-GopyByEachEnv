package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Go でのエラーハンドリングのセオリーは？
// ファイル操作系、より簡易な方法がある？　あまりにダラダラ。
// この事例ではゴルーチン使えないか・・・！？
// 複数ソースファイル化！

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadConfig() (m map[string]string) {
	var fp *os.File
	fp, err := os.OpenFile("config.txt", os.O_RDONLY, 0)
	handleErr(err)

	defer func() {
		fp.Close()
	}()

	m = map[string]string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
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

	return m
}

func CopyToEachDir(m map[string]string) {
	fromDir := m["fromDir"]
	fromFile := m["fromFile"]
	from := fromDir + fromFile
	toFiles := m["to"]
	toArray := strings.Split(toFiles, "$")

	for _, to := range toArray {
		src, err := os.Open(from)
		handleErr(err)
		defer src.Close()

		toSplit := strings.Split(to, "\t")
		dst, err := os.Create(toSplit[0])
		handleErr(err)
		defer dst.Close()

		fmt.Println("[COPY] " + from + " => " + toSplit[0])
		_, err = io.Copy(dst, src)
		handleErr(err)
	}
	fmt.Println("")
}

func ReplaceEachToFile(m map[string]string) {
	var fp *os.File
	var err error
	var fp2 *os.File
	var err2 error
	toFiles := m["to"]
	toArray := strings.Split(toFiles, "$")

	for _, to := range toArray {
		toSplit := strings.Split(to, "\t")
		fp, err = os.OpenFile(toSplit[0], os.O_APPEND, 0777)
		fp2, err2 = os.Create(toSplit[0] + ".tmp")
		handleErr(err)
		handleErr(err2)
		defer fp.Close()
		defer fp2.Close()
		scanner := bufio.NewScanner(fp)
		writer := bufio.NewWriter(fp2)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.Replace(line, toSplit[1], toSplit[2], -1)
			writer.WriteString(line + "\r\n")
		}
		writer.Flush()
		handleErr(scanner.Err())
	}
}

func RenameTmp(m map[string]string) {
	var err error
	toFiles := m["to"]
	toArray := strings.Split(toFiles, "$")

	for _, to := range toArray {
		toSplit := strings.Split(to, "\t")
		err = os.Rename(toSplit[0]+".tmp", toSplit[0])
		handleErr(err)
	}
}

func ModTime(m map[string]string) time.Time {
	fromDir := m["fromDir"]
	fromFile := m["fromFile"]
	from := fromDir + fromFile
	fp, err := os.Stat(from)
	handleErr(err)
	return fp.ModTime()
}

func main() {
	http.HandleFunc("/gopy/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start!")
		m := ReadConfig()

		var baseTime time.Time
		for {
			var nowTime = ModTime(m)
			if baseTime != nowTime {
				CopyToEachDir(m)
				ReplaceEachToFile(m)
				RenameTmp(m)
				baseTime = nowTime
			}

			time.Sleep(10 * time.Second)
		}
	})

	http.HandleFunc("/gopy/end", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("End!")
		os.Exit(0)
	})

	err := http.ListenAndServe(":7109", nil)
	handleErr(err)
}
