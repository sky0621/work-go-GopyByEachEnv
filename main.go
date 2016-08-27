package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func readConfig() {
	fmt.Println("read config start.")

	var fp *os.File
	fp, err := os.OpenFile("config.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		fp.Close()
	}()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/fs/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start!")
		readConfig()
		fmt.Println("after readConfig")
	})

	err := http.ListenAndServe(":7109", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
