package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var r io.Reader
	if os.Args[1] != "" {
		r = strings.NewReader(os.Args[1])
	} else {
		r = bufio.NewReader(os.Stdin)
	}
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	asText := string(bytes[:])
	asWords := strings.Split(asText, " ")
	for _, word := range asWords {
		fmt.Fprint(os.Stdout, word+" ")
		word = strings.TrimSpace(word)
		if res, err := exec.Command("play", word+".wav").Output(); err != nil {
			fmt.Println(res, err, word+".wav")
		}
	}
}
