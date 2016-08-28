package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
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
}

func ReplaceEachToFile(m map[string]string) {
	var fp *os.File
	var err error
	toFiles := m["to"]
	toArray := strings.Split(toFiles, "$")

	for _, to := range toArray {
		toSplit := strings.Split(to, "\t")
		fp, err = os.OpenFile(toSplit[0], os.O_WRONLY, 0)
		handleErr(err)
		defer fp.Close()
		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			line := scanner.Text()
			rep := regexp.MustCompile("`" + toSplit[1] + "`")
			line = rep.ReplaceAllString(line, toSplit[2])
		}
		handleErr(scanner.Err())
	}
}

func main() {
	http.HandleFunc("/fs/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start!")
		m := ReadConfig()
		CopyToEachDir(m)
		ReplaceEachToFile(m)
	})

	err := http.ListenAndServe(":7109", nil)
	handleErr(err)
}
