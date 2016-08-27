package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

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
	fmt.Println(toFiles)
	toArray := strings.Split(toFiles, "$")

	for _, to := range toArray {
		src, err := os.Open(from)
		handleErr(err)
		defer src.Close()

		dst, err := os.Create(to)
		handleErr(err)
		defer dst.Close()

		_, err = io.Copy(dst, src)
		handleErr(err)
	}
}

func main() {
	http.HandleFunc("/fs/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start!")
		m := ReadConfig()
		fmt.Println(m)
		CopyToEachDir(m)
	})

	err := http.ListenAndServe(":7109", nil)
	handleErr(err)
}
