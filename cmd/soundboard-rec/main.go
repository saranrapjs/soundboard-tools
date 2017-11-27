package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	survey "gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
)

func main() {

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lines := string(bytes[:])
	for _, word := range strings.Split(lines, "\n") {
		if _, err := os.Stat(word + ".wav"); !os.IsNotExist(err) {
			continue
		}
		if word != "" {
			recordWord(word)
		}
	}
}

func confirmWord(word string) (bool, error) {
	name := ""
	prompt := &survey.Input{
		Message: "Play recording [p], record again [r] or accept recording [y]",
	}
	err := survey.AskOne(prompt, &name, nil)
	var result bool
	switch strings.ToLower(name) {
	case "p":
		exec.Command("play", word+".wav").Run()
		return confirmWord(word)
	case "y":
		result = true
	}
	return result, err
}

const timeout = 10 * time.Second

var silenceArgs = strings.Split("silence 1 0.1 1% reverse silence 1 0.1 1% reverse", " ")

func trimFile(word string) {
	args := append([]string{word + ".wav", word + ".temp.wav"}, silenceArgs...)
	exec.Command("sox", args...).Run()
	exec.Command("mv", word+".temp.wav", word+".wav").Run()
}

func recordWord(word string) bool {
	cmd := exec.Command("rec", word+".wav")
	fmt.Println("Recording '" + word + "', press enter to stop recording")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan error)
	end := make(chan struct{})
	go func() {
		done <- cmd.Wait()
	}()
	go func() {
		fmt.Scanln()
		close(end)
	}()
	select {
	case <-time.After(timeout):
	case <-end:
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill: ", err)
		}
	case err := <-done:
		if err != nil {
			log.Fatal("recording errored: ", err)
		}
	}
	trimFile(word)

	ok, err := confirmWord(word)
	switch {
	case err == terminal.InterruptErr:
		return false
	case !ok:
		return recordWord(word)
	default:
		return true
	}
}
